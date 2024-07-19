package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"strings"
	"unicode/utf8"
)

type baseController struct {
	conn    net.Conn
	keyGUID []byte
}

var closeCodes map[int]string = map[int]string{
	1000: "NormalError",
	1001: "GoingAwayError",
	1002: "ProtocolError",
	1003: "UnknownType",
	1007: "TypeError",
	1008: "PolicyError",
	1009: "MessageTooLargeError",
	1010: "ExtensionError",
	1011: "UnexpectedError",
}

var BaseController = &baseController{
	keyGUIF: []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"),
}

func (ctr *baseController) computeAcceptKey(challengeKey string) string {
	h := sha1.New()
	h.Write([]byte(challengeKey))
	h.Write(ctr.keyGUID)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (ctr *baseController) Init(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	var err error
	ctr.conn, _, err = hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ctr.conn.Close()
		return
	}
	challengeKey := ctr.computeAcceptKey(r.Header.Get("Sec-Websocket-Key"))
	//建立连接时的固定写法
	lines := []string{
		"HTTP/1.1 101 Web Socket Protocol Handshake",
		"Server: go/echoserver",
		"Upgrade: WebSocket",
		"Connection: Upgrade",
		"Sec-WebSocket-Accept: " + challengeKey,
		"",
		"", // required for extra CRLF
	}

	if _, err = ctr.conn.Write([]byte(strings.Join(lines, "\r\n"))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		ctr.conn.Close()
		return
	}
	go ctr.Recv()
}

func (ctr *baseController) Recv() {
	for {
		fr, err := ctr.read()
		if err != nil {
			return
		}
		switch fr.opcode {
		case 8: //close
			ctr.conn.Close()
			return
		case 9: //Ping
			fr.pong()
			fallthrough
		case 0, 1, 2: // continuation, text, binary
			log.Println("Recv:", fr.test())
			if err = ctr.SendText("aaaaaaaaa"); err != nil {
				log.Println("Error sending", err)
				return
			}
		}
	}
}

func (ctr *baseController) read() (*frame, error) {
	fr := &frame{}
	head, err := ctr.slice(2)
	if err != nil {
		return fr, err
	}
	fr.isFragment = (head[0] & 0x80) == 0x00
	fr.opcode = head[0] & 0x0F
	fr.reserved = (head[0] & 0x70)
	fr.ismasked = (head[1] & 0x80) == 0x80
	length := uint64(head[1] & 0x7F)
	if length == 126 {
		data, err := ctr.slice(2)
		if err != nil {
			return fr, err
		}
		length = uint64(binary.BigEndian.Uint16(data))
	} else if length == 127 {
		data, err := ctr.slice(8)
		if err != nil {
			return fr, err
		}
		length = binary.BigEndian.Uint64(data)
	}
	mask, err := ctr.slice(4)
	if err != nil {
		return fr, err
	}
	fr.length = length
	payload, err := ctr.slice(int(length))
	if err != nil {
		return fr, err
	}
	for i := uint64(0); i < length; i++ {
		payload[i] ^= mask[i%4]
	}
	fr.payload = payload
	err = ctr.validate(fr)
	return fr, err
}

func (ctr *baseController) validate(fr *frame) error {
	if !fr.ismasked {
		return errors.New("protocol error: unmasked client frame")
	}
	if fr.isControl() && (fr.length > 125 || fr.isFragment) {
		return errors.New("protocol error: all control frames MUST have a payload length of 125 bytes or less and MUST NOT be fragmented")
	}
	if fr.hasReservedOpcode() {
		return errors.New("protocol error: opcode " + fmt.Sprintf("%x", fr.opcode) + " is reserved")
	}
	if fr.reserved > 0 {
		return errors.New("protocol error: RSV " + fmt.Sprintf("%x", fr.reserved) + " is reserved")
	}
	if fr.opcode == 1 && !fr.isFragment && !utf8.Valid(fr.payload) {
		return errors.New("wrong code: invalid UTF-8 text message ")
	}
	if fr.opcode == 8 {
		if fr.length >= 2 {
			code := binary.BigEndian.Uint16(fr.payload[:2])
			reason := utf8.Valid(fr.payload[2:])
			if code >= 5000 || (code < 3000 && closeCodes[int(code)] == "") {
				return errors.New(closeCodes[1002] + " Wrong Code")
			}
			if fr.length > 2 && !reason {
				return errors.New(closeCodes[1007] + " invalid UTF-8 reason message")
			}
		} else if fr.length != 0 {
			return errors.New(closeCodes[1002] + " Wrong Code")
		}
	}
	return nil
}
func (ctr *baseController) slice(size int) ([]byte, error) {
	data := make([]byte, 0)
	for {
		if len(data) == size {
			break
		}
		sz := 4096
		remaining := size - len(data)
		if sz > remaining {
			sz = remaining
		}
		temp := make([]byte, sz)
		n, err := ctr.conn.Read(temp)
		if err != nil && err != io.EOF {
			return data, err
		}
		data = append(data, temp[:n]...)
	}
	return data, nil
}

func (ctr *baseController) SendText(txt string) error {
	fr := &frame{
		isFragment: false,
		opcode:     1,
		payloaF:    []byte(txt),
	}
	fr.length = uint64(len(fr.payload))
	return ctr.Send(fr)
}

func (ctr *baseController) Send(fr *frame) error {
	data := make([]byte, 2)
	data[0] = 0x80 | fr.opcode
	if fr.isFragment {
		data[0] &= 0x7F
	}
	if fr.length <= 125 {
		data[1] = byte(fr.length)
		data = append(data, fr.payload...)
	} else if fr.length > 125 && float64(fr.length) < math.Pow(2, 16) {
		data[1] = byte(126)
		size := make([]byte, 2)
		binary.BigEndian.PutUint16(size, uint16(fr.length))
		data = append(data, size...)
		data = append(data, fr.payload...)
	} else if float64(fr.length) >= math.Pow(2, 16) {
		data[1] = byte(127)
		size := make([]byte, 8)
		binary.BigEndian.PutUint16(size, uint16(fr.length))
		data = append(data, size...)
		data = append(data, fr.payload...)
	}
	if _, err := ctr.conn.Write(data); err != nil {
		return err
	}
	return nil
}

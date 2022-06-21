package controllers

import (
	"bytes"
	"encoding/binary"
)

type frame struct {
	isFragment bool
	/*	0 continuation
		1 text
		2 binary
		8 close
		9 ping
		10 pong
	*/
	opcode     byte
	reserved   byte
	ismasked   bool
	length     uint64
	payload    []byte
}

func (f *frame) pong() {
	f.opcode = 10
}

func (f *frame) test() string {
	return string(f.payload)
}

func (f *frame) isControl() bool {
	return f.opcode&0x08 == 0x08
}

func (f *frame) hasReservedOpcode() bool {
	return f.opcode > 10 || (f.opcode >= 3 && f.opcode <= 7)
}

func (f *frame) closeCode() (code uint16) {
	binary.Read(bytes.NewReader(f.payload), binary.BigEndian, &code)
	return
}
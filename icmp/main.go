package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout int
	size    int
	count   int
)

func doflag() {
	flag.IntVar(&timeout, "w", 1000, "设置超时时间,单位毫秒")
	flag.IntVar(&size, "l", 32, "设置报文大小")
	flag.IntVar(&count, "n", 4, "设置报文发送次数")
	flag.Parse()
}

type ICMP struct {
	Type     uint8  //类型
	Code     uint8  //代码
	CheckSum uint16 //校验和
	ID       uint16 //标识
	Sequence uint16 //序列号
}

/*
	计算校验和

将数据每16位相加，得到一个32位的数
再将32位的数的高16位与低16位不断相加,直到高16位全为0
最后结果取反
*/
func checkSum(data []byte) uint16 {
	size := len(data)
	i := 0
	sum := uint32(0)
	for size > 1 {
		sum += (uint32(data[i]) << 8) + uint32(data[i+1])
		i += 2
		size -= 2
	}
	if size == 1 {
		sum += uint32(data[i])
	}
	low := uint16(sum)
	high := uint16(sum >> 16)
	for high != 0 {
		sum = uint32(low) + uint32(high)
		low = uint16(sum)
		high = uint16(sum >> 16)
	}
	return ^low
}

func main() {
	doflag()
	// addr需要是最后一个参数
	addr := os.Args[len(os.Args)-1]
	conn, err := net.Dial("ip4:icmp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	icmp := &ICMP{
		Type:     8,
		Code:     0,
		CheckSum: 0,
		IF:       1,
		Sequence: 1,
	}
	data := make([]byte, size)
	buffer := bytes.Buffer{}
	binary.Write(&buffer, binary.BigEndian, icmp)
	buffer.Write(data)
	data = buffer.Bytes()
	fmt.Println(data)
	sum := checkSum(data)
	data[3] = byte(sum)      //253
	data[2] = byte(sum >> 8) //247
	fmt.Println(data)

	fmt.Printf("正在Ping %s [%s] 具有%d字节数据: \n", addr, conn.RemoteAddr(), size)
	for x := 0; x < count; x++ {
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		t1 := time.Now().UnixMilli()
		_, err = conn.Write(data)
		if err != nil {
			log.Println("1", err)
			continue
		}
		buff := make([]byte, 65535)
		n, err := conn.Read(buff)
		if err != nil {
			log.Println("2", err)
			continue
		}
		t2 := time.Now().UnixMilli()
		fmt.Printf(" 来自 %d.%d.%d.%d 的恢复：字节=%d, 时间=%dms, TTL=%d \n", buff[12], buff[13], buff[14], buff[15], n-28, t2-t1, buff[8])
		time.Sleep(time.Second)
	}
}

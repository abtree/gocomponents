package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type World struct {
}

//rpc函数必须是两个参数 传入参数 和 传出参数
func (this *World) HelloWorld(name *string, resp *string) error {
	*resp = *name + ": Server say hello!"
	return nil
}

func main() {
	// 注册服务
	err := rpc.Register(new(World))
	if err != nil {
		log.Fatalln("rpc register error: ", err.Error())
	}
	// 设置监听
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalln("net listen error: ", err.Error())
	}
	defer listener.Close()
	// 建立连接
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("net Accept error: ", err.Error())
	}
	defer conn.Close()
	// 绑定服务
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}

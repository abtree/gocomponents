package main

import (
	"log"
	"net"
	"net/rpc"
)

//接口标明每个rpc函数需要的参数和返回值
type WorldInterface interface {
	HelloWorld(string, *string) error
}

//World实现WorldInterface接口
type World struct {
}

//rpc函数必须是两个参数 传入参数 和 传出参数
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + ": Server say hello!"
	return nil
}

//封装这个函数 是为了通过接口检查rpc函数是否正确定义
func RegisterServer(name string, p WorldInterface) {
	err := rpc.RegisterName(name, p)
	if err != nil {
		log.Fatalln("rpc register error: ", err.Error())
	}
}

func main() {
	// 注册服务
	RegisterServer("hello", new(World))
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
	rpc.ServeConn(conn)
}

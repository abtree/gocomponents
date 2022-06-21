package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//接口标明每个rpc函数需要的参数和返回值
type WorldInterface interface {
	HelloWorld(string, *string) error
}

//封装客户端，使其能在编译时检查参数错误
type MyClient struct {
	c *rpc.Client
}

func (this *MyClient) HelloWorld(a string, b *string) error {
	/*调用远程函数
	  参数1： 服务名.函数名
	  参数2： 传入参数
	  参数3： 传出参数
	  go语音默认是使用gob序列化，在与其它语音通信时，会出现乱码问题，
	  如果需要与其它语音通信，可以使用grpc或jsonrpc
	*/
	return this.c.Call("hello.HelloWorld", a, b)
}

func InitClient(conn *rpc.Client) *MyClient {
	return &MyClient{c: conn}
}

func main() {
	//用rpc连接服务器
	conn, err := rpc.Dial("tcp", ":8082")
	if err != nil {
		log.Fatalln("net Dial error: ", err.Error())
	}
	defer conn.Close()

	var resp string            //用于接收传出参数
	client := InitClient(conn) //这里封装是为了在编译时就检查错误
	err = client.HelloWorld("input params", &resp)
	//err = conn.Call("hello.HelloWorld", "input params", &resp)
	if err != nil {
		log.Fatalln("rpc call error: ", err.Error())
	}
	fmt.Println(resp)
}

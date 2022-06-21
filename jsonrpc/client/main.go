package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//用rpc连接服务器
	conn, err := net.Dial("tcp", ":8082")
	if err != nil {
		log.Fatalln("net Dial error: ", err.Error())
	}
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var resp string //用于接收传出参数
	/*调用远程函数
	  参数1： 服务名.函数名
	  参数2： 传入参数
	  参数3： 传出参数
	  此处参数会以json格式传输：大致如下
	{"method":"World.HelloWorld", "params":["input params"], "id":0}
	  返回也是以json格式传输
	{"id":0,"result":"input params: Server say hello!", "error":null}
	*/
	err = client.Call("World.HelloWorld", "input params", &resp)
	if err != nil {
		log.Fatalln("rpc call error: ", err.Error())
	}
	fmt.Println(resp)
}

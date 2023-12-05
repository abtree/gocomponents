package main

/*
 Token认证，可用多种认证方式，如google的Auth2
 当前实现：用户名密码认证
 注意：这个可以和证书认证同时使用
*/

import (
	"components/grpc/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	mtdata := metadata.New(map[string]string{"user": "admin", "passwd": "admin"})
	ctx := metadata.NewOutgoingContext(context.Background(), mtdata)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

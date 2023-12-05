package main

/*
客户端实现单向认证
*/

import (
	"components/grpc/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	file, err := credentials.NewClientTLSFromFile("certs/server.pem", "localhost")
	if err != nil {
		log.Panicln(err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(file))
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

package main

/*
客户端实现双向认证
*/

import (
	"components/grpc/pb"
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Panicln(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("certs/ca.pem")
	if err != nil {
		log.Panicln(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //客户端的证书
		ServerName:   "localhost",
		RootCAs:      certPool, //用于颁发证书的ca，客户端与服务器必须相同
	})
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
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

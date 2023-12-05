package main

/*
服务器端实现双向认证
*/

import (
	"components/grpc/pb"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(req.Name)
	back := &pb.HelloReply{Message: req.Name}
	return back, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalln(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("certs/ca.pem")
	if err != nil {
		log.Fatalln(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务器的证书
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool, //用于颁发证书的ca，客户端与服务器必须相同
	})
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

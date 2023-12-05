package main

/*
服务器端实现单向认证
*/

import (
	"components/grpc/pb"
	"context"
	"fmt"
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
	file, err := credentials.NewServerTLSFromFile("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer(grpc.Creds(file))
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"components/grpc/pb"
	"io"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(req.Name)
	back := &pb.HelloReply{Message: req.Name}
	return back, nil
}

func (*server) ListHello(req *pb.HelloRequest, srv pb.Greeter_ListHelloServer) error {
	fmt.Println(req.Name)
	for i := 0; i < 3; i++ {
		back := &pb.HelloReply{Message: req.Name}
		srv.Send(back)
	}
	return nil
}
func (*server) RecordHello(srv pb.Greeter_RecordHelloServer) error {
	for {
		recv, err := srv.Recv()
		if err == io.EOF {
			srv.SendAndClose(&pb.HelloReply{Message: "Hello"})
			return err
		}
		if err != nil {
			log.Fatalf("%v", err)
			return err
		}
		log.Printf("%v", recv.Name)
	}
	return nil
}
func (*server) ChatHello(srv pb.Greeter_ChatHelloServer) error {
	md, ok := metadata.FromIncomingContext(srv.Context())
	if !ok {
		log.Fatalf("context error")
		return errors.New("context error")
	}
	if len(md["init"]) == 0 {
		log.Fatalf("context error empty")
		return errors.New("context error empty")
	}
	log.Printf("%v", md["init"])

	for {
		r, err := srv.Recv()
		if err == io.EOF {
			log.Printf("client closed %v", err)
			return err
		}
		if err != nil {
			log.Fatalf("error %v", err)
			return err
		}
		srv.Send(&pb.HelloReply{Message: r.Name})
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

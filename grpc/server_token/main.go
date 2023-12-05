package main

/*
 Token认证，可用多种认证方式，如google的Auth2
 当前实现：用户名密码认证
 注意：这个可以和证书认证同时使用
*/

import (
	"components/grpc/pb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	//定义一个拦截器
	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = Auth(ctx)
		if err != nil {
			log.Panicln(err)
		}
		return handler(ctx, req)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var user string
	var passwd string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["passwd"]; ok {
		passwd = val[0]
	}
	if user != "admin" || passwd != "admin" {
		return fmt.Errorf("user or password error")
	}
	return nil
}

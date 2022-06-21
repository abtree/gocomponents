package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"components/grpc/pb"
	"log"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	mtdata := metadata.New(map[string]string{"init": "Hallo"})
	ctx := metadata.NewOutgoingContext(context.Background(), mtdata)

	//单通道
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.GetMessage())

	//服务器流
	//stream, err := c.ListHello(ctx, &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//for {
	//	in, err := stream.Recv()
	//	if err == io.EOF {
	//		log.Printf("Close %v", err)
	//		return
	//	}
	//	if err != nil {
	//		log.Fatalf("Fatal %d", err)
	//		return
	//	}
	//	fmt.Printf("%v", in.Message)
	//}

	//客户端流
	//stream, err := c.RecordHello(ctx)
	//for i := 0; i < 3; i++ {
	//	stream.Send(&pb.HelloRequest{Name: name})
	//}
	//r, err := stream.CloseAndRecv()
	//log.Printf("Greeting: %s", r.GetMessage())

	//双向流
	stream, err := c.ChatHello(ctx)
	for i := 0; i < 3; i++ {
		stream.Send(&pb.HelloRequest{Name: name})
		r, err := stream.Recv()
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
	stream.CloseSend()
}

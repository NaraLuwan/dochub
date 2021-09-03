package main

import (
	"context"
	"fmt"
	helloworld "go.opencensus.io/examples/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	// 注册grpcurl所需的reflection服务
	reflection.Register(server)
	// 注册业务服务
	helloworld.RegisterGreeterServer(server, &greeter{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type greeter struct {
}

func (*greeter) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println(req)
	reply := &helloworld.HelloReply{Message: "hello"}
	return reply, nil
}

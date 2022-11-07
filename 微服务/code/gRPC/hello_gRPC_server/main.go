package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello_gRPC_server/proto"
	"net"
)

// hello server
type server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Reply: "hello " + in.Name}, nil
}

func main() {
	// 监听本地的8972端口
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                          // 创建gRPC服务器
	proto.RegisterHelloServiceServer(s, &server{}) // 在gRPC服务端注册服务
	//	启动服务
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

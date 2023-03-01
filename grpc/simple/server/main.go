package main

import (
	"context"
	"fmt"
	"github.com/playmood/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

//type HelloServiceServer interface {
//	Hello(context.Context, *Request) (*Response, error)
//	mustEmbedUnimplementedHelloServiceServer()
//}

// HelloServiceServer must be embedded to have forward compatible implementations.
type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello, %s", req.Value)}, nil
}

func main() {
	// s grpc.ServiceRegistrar, srv HelloServiceServer
	// srv HelloServiceServer 实现类
	server := grpc.NewServer()

	// 把实现类注册给grpc server
	pb.RegisterHelloServiceServer(server, new(HelloServiceServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	log.Printf("grpc listen addr: 127.0.0.1:1234")
	// 监听socket, HTTP2内置
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

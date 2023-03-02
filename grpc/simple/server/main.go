package main

import (
	"context"
	"fmt"
	server2 "github.com/playmood/rpc/grpc/middleware/server"
	"github.com/playmood/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"io"
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

func (s *HelloServiceServer) Channel(stream pb.HelloService_ChannelServer) error {
	// 处理单个连接 放在goroutine中跑的 天生高并发
	for {
		// 接收请求
		req, err := stream.Recv()
		if err != nil {
			// 当前客户端退出
			if err == io.EOF {
				log.Println("client closed")
				return nil
			}
			return err
		}

		resp := &pb.Response{Value: fmt.Sprintf("hello %s", req.Value)}
		// 响应请求
		err = stream.Send(resp)
		if err != nil {
			if err == io.EOF {
				log.Println("client closed")
				return nil
			}
			return err
		}
	}
}

func main() {
	// s grpc.ServiceRegistrar, srv HelloServiceServer
	// srv HelloServiceServer 实现类
	reqAuth := server2.NewGrpcAuthUnaryServerInterceptor()
	streamAuth := server2.NewGrpcAuthStreamServerInterceptor()
	server := grpc.NewServer(grpc.UnaryInterceptor(reqAuth), grpc.StreamInterceptor(streamAuth))

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

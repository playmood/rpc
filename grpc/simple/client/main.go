package main

import (
	"context"
	"fmt"
	"github.com/playmood/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	// grpc为我们生成一个客户端调用服务端的SDK
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)

	// request response
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "Alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// stream
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 客户端发送放进单独的goroutine处理
	go func() {
		for {
			if err := stream.Send(&pb.Request{Value: "dd"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 主循环接收服务端响应数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}

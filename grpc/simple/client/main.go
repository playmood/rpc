package main

import (
	"context"
	"fmt"
	"github.com/playmood/rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {
	// grpc为我们生成一个客户端调用服务端的SDK
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "Alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

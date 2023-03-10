package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/playmood/rpc/protobuf-rpc/service"
)

// 约束客户端
var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	// c, err := rpc.Dial(network, address)
	// if err != nil {
	// 	return nil, err
	// }

	// 建立链接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	// 采用Json编解码的客户端
	c := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(req *service.Request, resp *service.Response) error {
	return p.Client.Call(service.SERVICE_NAME+".Hello", req, resp)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	resp := &service.Response{}
	err = client.Hello(&service.Request{Value: "hello"}, resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

package main

import (
	"fmt"
	"github.com/playmood/rpc/rpc-TCP-json/service"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 约束客户端接口实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// 客户端采用json格式来编解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, response *string) error {

	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
}

func main() {
	var resp string
	client, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	if err := client.Hello("alice 01", &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

package main

import (
	"fmt"
	"github.com/playmood/rpc/rpc-interface/service"
	"net/rpc"
)

// 约束客户端接口实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		return nil, err
	}
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
	if err := client.Hello("Bob", &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

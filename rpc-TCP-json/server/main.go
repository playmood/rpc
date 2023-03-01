package main

import (
	"fmt"
	"github.com/playmood/rpc/rpc-TCP-json/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 约束接口实现
var _ service.HelloService = (*HelloService)(nil)

type HelloService struct {
}

// 注册rpc服务的统一签名
func (s *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil
}

func main() {
	// rpc对外暴露的对象注册到rpc框架中
	rpc.RegisterName("HelloService", &HelloService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// 每个客户端单独启用一个routine处理
		// 采用json来进行编解码，类似于json.Unmarshal和 Marshal
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

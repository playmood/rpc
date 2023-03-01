package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

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
		go rpc.ServeConn(conn)
	}
}

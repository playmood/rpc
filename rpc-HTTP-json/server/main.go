package main

import (
	"fmt"
	"github.com/playmood/rpc/rpc-HTTP-json/service"
	"io"
	"net/http"
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

func NewRPCReadWriteCloseFromHTTP(writer http.ResponseWriter, request *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{writer, request.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func main() {
	// rpc对外暴露的对象注册到rpc框架中
	rpc.RegisterName("HelloService", &HelloService{})

	// rpc服务架设在/jsonrpc
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		conn := NewRPCReadWriteCloseFromHTTP(writer, request)
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}

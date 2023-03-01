# 生成代码

```shell
> cd github.com/playmood/rpc/grpc/simple/server
# github.com/playmood/rpc/grpc/simple/server
> protoc -I=./pb/ --go_out=. --go_opt=module="github.com/playmood/rpc/grpc/simple/server" ./pb/hello.proto
```

# rpc接口定义 protobuf的代码生成
```shell
> protoc -I=./pb/ --go_out=. --go_opt=module="github.com/playmood/rpc/grpc/simple/server" --go-grpc_out=. --go-grpc_opt=module="github.com/playmood/rpc/grpc/simple/server" ./pb/hello.proto
```
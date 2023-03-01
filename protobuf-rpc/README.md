# protobuf 代码生成

```shell
protoc -I=./service/pb/ --go_out=./service/ --go_opt=module="github.com/playmood/rpc/protobuf-rpc/service" hello.proto
```
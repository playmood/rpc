# 如何生成代码

```shell
# 在protobuf文件夹下
protoc -I=./pb/ --go_out=./pb/ --go_opt=module="github.com/playmood/rpc/protobuf/pb" hello.proto
```
grpcurl是一个命令行工具，使用它可以在命令行中访问gRPC服务，就像使用curl访问http服务一样

## 准备
在gRPC服务中注册reflection服务，gRPC服务是使用Protobuf(PB)协议的，而PB提供了在运行时获取Proto定义信息的反射功能。grpc-go中的"google.golang.org/grpc/reflection"包就对这个反射功能提供了支持。

这里以grpc-go官方的helloword为例，代码结构如下:
```text
grpc-hello
├── go.mod
├── go.sum
├── main.go
└── proto
    ├── doc.go
    ├── helloworld.pb.go
    └── helloworld.proto
```
proto定义详见：[helloworld.proto:](https://github.com/NaraLuwan/dochub/blob/master/GRpc/helloworld.proto)

注册grpc服务详见：[main.go](https://github.com/NaraLuwan/dochub/blob/master/GRpc/main.go)，在main.go的第19行，使用reflection.Register(server)注册了reflection服务

## grpcurl的安装和使用
在Mac OS下安装grpcurl:
```shell script
brew install grpcurl
```
grpcurl使用如下示例如下:

查看服务列表:
```shell script
grpcurl -plaintext 127.0.0.1:8080 list
```
```text
grpc.reflection.v1alpha.ServerReflection
proto.Greeter
```

查看某个服务的方法列表:
```shell script
grpcurl -plaintext 127.0.0.1:8080 list proto.Greeter
```
```text
proto.Greeter.SayHello
```

查看方法定义:
```shell script
grpcurl -plaintext 127.0.0.1:8080 describe proto.Greeter.SayHello
```
```text
proto.Greeter.SayHello is a method:
rpc SayHello ( .proto.HelloRequest ) returns ( .proto.HelloReply );
```

查看请求参数:
```shell script
grpcurl -plaintext 127.0.0.1:8080 describe proto.HelloRequest
```
```text
proto.HelloRequest is a message:
message HelloRequest {
  string name = 1;
}
```

调用服务，参数传json即可:
```shell script
grpcurl -d '{"name": "abc"}' -plaintext 127.0.0.1:8080  proto.Greeter.SayHello
```
{
  "message": "hello"
}

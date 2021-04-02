# Briscola

[gobyexample http-server](https://gobyexample.com/http-servers)
[go-micro examples](https://github.com/asim/go-micro/tree/master/examples)
[go-micro proto example](https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro)

[grpc docs](https://grpc.io/docs/languages/go/quickstart/)
[protobuf sytnax](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec#syntax)
[protobuf gen guide](https://developers.google.com/protocol-buffers/docs/reference/go-generated)

```shell
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
cd protoc-3.15.6-linux-x86_64/
go install google.golang.org/protobuf/cmd/protoc-gen-go
export PATH #if protoc-gen-go is not inside path
protoc --proto_path=dom/briscola/ --go_out=dom/briscola briscola.proto
```
#!/bin/sh
set -ev
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc dom/briscola/pb/*.proto --go_out=plugins=grpc:.
echo
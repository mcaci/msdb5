#!/bin/sh
set -ev
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc dom/briscola/pb/points.proto --go_out=plugins=grpc:.
protoc dom/briscola/v1/pb/*.proto --go_out=plugins=grpc:.
echo
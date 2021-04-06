package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mcaci/msdb5/v2/dom/briscola/srv"
	"github.com/mcaci/msdb5/v2/pb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "port where briscola services are listenning")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBriscolaServer(grpcServer, srv.NewServer())
	grpcServer.Serve(lis)
}

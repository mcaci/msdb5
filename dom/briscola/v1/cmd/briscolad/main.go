package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mcaci/msdb5/dom/briscola/v1/pb"
	briscola "github.com/mcaci/msdb5/dom/briscola/v1/server/grpc"
	briscolahttp "github.com/mcaci/msdb5/dom/briscola/v1/server/http"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	"google.golang.org/grpc"
)

func main() {

	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := serv.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	pointsEndpoint := serv.MakePointsEndpoint(srv)
	countEndpoint := serv.MakeCountEndpoint(srv)
	compareEndpoint := serv.MakeCompareEndpoint(srv)
	endpoints := serv.Endpoints{
		CardPointsEndpoint:  pointsEndpoint,
		PointCountEndpoint:  countEndpoint,
		CardCompareEndpoint: compareEndpoint,
	}

	// start HTTP server
	go func() {
		log.Println("http:", *httpAddr)
		handler := briscolahttp.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// start gRPC server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Println("grpc:", *gRPCAddr)
		handler := briscola.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterBriscolaServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	log.Fatalln(<-errChan)
}

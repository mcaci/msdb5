package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	grpcclient "github.com/mcaci/msdb5/dom/briscola/v1/client/grpc"
	serv "github.com/mcaci/msdb5/dom/briscola/v1/service"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("addr", ":8081", "gRPC address")
	)
	flag.Parse()
	ctx := context.Background()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()
	pointsService := grpcclient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "points":
		var number string
		number, args = pop(args)
		n, _ := strconv.Atoi(number)
		points(ctx, pointsService, uint32(n))
	default:
		log.Fatalln("unknown command", cmd)
	}
}

func points(ctx context.Context, service serv.Service, number uint32) {
	h, err := service.CardPoints(ctx, number)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(h)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

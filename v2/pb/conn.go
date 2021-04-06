package pb

import (
	"flag"
	"log"

	grpc "google.golang.org/grpc"
)

func Conn() *grpc.ClientConn {
	const serverAddr = "localhost:8081"
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Println("error found", err)
	}
	return conn
}

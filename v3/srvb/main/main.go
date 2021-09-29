package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v3/srvb"
)

func main() {
	port := flag.String("port", "8080", "port where the briscola server is listening to")
	flag.Parse()
	err := http.ListenAndServe(":"+*port, srvb.Handler())
	if err != nil {
		log.Fatal(err)
	}
}

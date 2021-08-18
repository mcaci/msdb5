package main

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/srvb"
)

func main() {
	err := http.ListenAndServe(":8080", srvb.Handler())
	if err != nil {
		log.Fatal(err)
	}
}

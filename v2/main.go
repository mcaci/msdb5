package main

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/srvb"
)

func main() {
	http.HandleFunc("/create", srvb.Create)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

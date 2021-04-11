package main

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/frw/srv"
)

func main() {
	http.HandleFunc("/", srv.Msdb5)
	http.HandleFunc("/start/", srv.Start)
	http.HandleFunc("/draw/", srv.Draw)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

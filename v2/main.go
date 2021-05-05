package main

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/frw/srv"
)

func main() {
	http.HandleFunc("/", srv.Home)
	http.HandleFunc("/start/", srv.Start)
	http.HandleFunc("/play/", srv.Play)
	http.HandleFunc("/refresh/", srv.Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

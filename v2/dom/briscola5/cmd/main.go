package main

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/dom/briscola5/srv"
)

func main() {
	http.HandleFunc("/cmp", srv.Cmp)
	http.HandleFunc("/cmpandset", srv.CmpAndSet)

	log.Fatal(http.ListenAndServe(":8082", nil))
}

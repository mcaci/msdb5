package main

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/dom/briscola5/cmd/mdw"
)

func main() {
	http.HandleFunc("/cmp", mdw.Cmp)
	http.HandleFunc("/cmpandset", mdw.CmpAndSet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

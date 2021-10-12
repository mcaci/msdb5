package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/pong", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("pong, send to ping")
		time.AfterFunc(time.Second, func() { http.Get("http://localhost:8082/ping") })
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

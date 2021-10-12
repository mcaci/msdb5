package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("ping, send to pong")
		time.AfterFunc(time.Second, func() { http.Get("http://localhost:8081/pong") })
	})
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

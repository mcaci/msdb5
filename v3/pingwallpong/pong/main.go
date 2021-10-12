package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/pong", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("pong")
		time.AfterFunc(time.Second, func() {
			_, err := http.Post("http://localhost:8083/wall", "application/json", strings.NewReader(`{"from":"pong"}`))
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		})
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

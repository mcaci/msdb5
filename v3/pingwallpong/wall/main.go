package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/wall", func(rw http.ResponseWriter, r *http.Request) {
		var f struct {
			From string `json:"from"`
		}
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&f)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		switch f.From {
		case "ping":
			fmt.Println("hit by", f.From, "send to pong")
			time.AfterFunc(time.Second, func() {
				_, err := http.Get("http://localhost:8081/pong")
				if err != nil {
					tick := time.NewTicker(time.Second)
					go time.AfterFunc(10*time.Second, tick.Stop)
					for range tick.C {
						fmt.Println("target not answering, retrying")
						_, err = http.Get("http://localhost:8081/pong")
						if err != nil {
							continue
						}
						tick.Stop()
						break
					}
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			})
		case "pong":
			fmt.Println("hit by", f.From, "send to ping")
			time.AfterFunc(time.Second, func() {
				_, err := http.Get("http://localhost:8082/ping")
				if err != nil {
					tick := time.NewTicker(time.Second)
					go time.AfterFunc(10*time.Second, tick.Stop)
					for range tick.C {
						fmt.Println("target not answering, retrying")
						_, err = http.Get("http://localhost:8082/ping")
						if err != nil {
							continue
						}
						tick.Stop()
						break
					}
					http.Error(rw, err.Error(), http.StatusInternalServerError)
				}
			})
		}
	})
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal(err)
	}
}

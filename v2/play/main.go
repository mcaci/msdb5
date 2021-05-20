package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	join := flag.Bool("join", false, "true for joining the game and false for creating it")
	flag.Parse()

	var n, t string
	switch *join {
	case true:
		n = "ai2"
		t = "join"
	case false:
		n = "ai1"
		t = "create"
	}
	resp, err := http.PostForm("http://localhost:8080/start/"+n, url.Values{"playername": []string{n}, "type": []string{t}, "gamename": []string{"default"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.Body)

	for i := 0; i < 25; i++ {
		resp, err := http.PostForm("http://localhost:8080/play/"+n, url.Values{"cardn": []string{"0"}})
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		log.Println(resp.Body)
		time.Sleep(2 * time.Second)
	}
}

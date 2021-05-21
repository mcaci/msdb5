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

	for {
		resp, err := http.PostForm("http://localhost:8080/play/"+n, url.Values{"cardn": []string{"0"}})
		if err != nil {
			log.Println(err)
			break
		}
		defer resp.Body.Close()
		p := []byte{}
		n, err := resp.Body.Read(p)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(n, p)
		if string(p) == "Game is over" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("Match ended")
}

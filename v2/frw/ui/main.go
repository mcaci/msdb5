package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/mcaci/ita-cards/set"
)

type Page struct {
	Title string
	Body  []byte
}

var (
	templates = template.Must(template.ParseFiles("player.html"))
)

func start(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "player.html", &Page{Title: "Player"})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func draw(w http.ResponseWriter, r *http.Request) {
	currentBody = append(currentBody, []byte(cards.Top().String()))
	p := &Page{Title: "Player", Body: bytes.Join(currentBody, []byte(","))}
	err := templates.ExecuteTemplate(w, "player.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var cards = set.Deck()
var currentBody = [][]byte{}

func main() {
	http.HandleFunc("/start/", start)
	http.HandleFunc("/draw/", draw)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

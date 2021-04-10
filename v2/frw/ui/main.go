package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/mcaci/ita-cards/set"
)

type Page struct {
	Title string
	Body  []byte
	Msg   []byte
}

var (
	files     = []string{"frw/ui/start.html", "frw/ui/player.html"}
	templates = template.Must(template.ParseFiles(files...))
)

func msdb5(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "start.html", &Page{Title: "Start", Body: []byte("test")})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	v := &Page{Title: "New Game"}
	name := r.Form["gamename"]
	switch r.Form["type"][0] {
	case "create":
		v.Msg = []byte(fmt.Sprintf("creating game %q", name))
	case "join":
		v.Msg = []byte(fmt.Sprintf("joining game %q", name))
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println(r.Form)
	err = templates.ExecuteTemplate(w, "start.html", v)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func draw(w http.ResponseWriter, r *http.Request) {
	currentBody = append(currentBody, []byte(cards.Top().String()))
	p := &Page{Title: "Player", Body: bytes.Join(currentBody, []byte(", "))}
	err := templates.ExecuteTemplate(w, "player.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var cards = set.Deck()
var currentBody = [][]byte{}

func main() {
	http.HandleFunc("/", msdb5)
	http.HandleFunc("/start/", start)
	http.HandleFunc("/draw/", draw)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package srv

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/game"
)

var (
	files     = []string{"assets/start.html", "assets/player.html"}
	templates = template.Must(template.ParseFiles(files...))

	n string
	g *game.Game

	cards       = set.Deck()
	currentBody = [][]byte{}
)

type Page struct {
	Title string
	Body  []byte
	Msg   []byte
}

func Msdb5(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "start.html", &Page{Title: "Start", Body: []byte("test")})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func Start(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	v := &Page{Title: "New Game"}
	name := r.Form["gamename"][0]
	switch r.Form["type"][0] {
	case "create":
		if g != nil {
			log.Printf("Game with name %q already exists, cannot creata game %q too", n, name)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		g = game.NewGame(&game.Options{WithSide: true})
		n = name
		v.Msg = []byte(fmt.Sprintf("new game created with name %q", name))
		log.Printf("Game created with name %q", name)
	case "join":
		if g == nil {
			log.Println("No games existing yet", name)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if name != n {
			log.Printf("game %s not found", name)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		v.Msg = []byte(fmt.Sprintf("joining game %q", name))
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "Did not understand the action", http.StatusInternalServerError)
		return
	}
	log.Println(r.Form)
	err = templates.ExecuteTemplate(w, "start.html", v)
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func Draw(w http.ResponseWriter, r *http.Request) {
	currentBody = append(currentBody, []byte(cards.Top().String()))
	p := &Page{Title: "Player", Body: bytes.Join(currentBody, []byte(", "))}
	err := templates.ExecuteTemplate(w, "player.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

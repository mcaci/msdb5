package srv

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/frw/session"
)

var (
	s    = session.Briscola{Deck: set.Deck()}
	s5   = session.Briscola5{}
	game = template.Must(template.ParseFiles("assets/game.html"))
)

func Start(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	playername := r.Form["playername"][0]
	gamename := r.Form["gamename"][0]
	var body []byte
	switch r.Form["type"][0] {
	case "create":
		if s.Game != nil && s.Game.Started(gamename) {
			log.Print("another game already exists, cannot create more than 1")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		s.Game = briscola.NewGame(&briscola.Options{
			WithSide: true,
			WithName: gamename,
		})
		err := briscola.Register(playername, s.Game)
		if err != nil {
			log.Print("registration error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		body = []byte(fmt.Sprintf("new game created with gamename %q by player %q", gamename, playername))
		log.Print(string(body))
	case "join":
		if s.Game == nil {
			log.Print("no games created yet")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if !s.Game.Started(gamename) {
			log.Printf("game %s not found", gamename)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		err := briscola.Register(playername, s.Game)
		if err != nil {
			log.Print("registration error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		body = []byte(fmt.Sprintf("player %q joining game %q", playername, gamename))
		log.Print(string(body))
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "did not understand the action", http.StatusInternalServerError)
		return
	}
	log.Print(s.Game)
	err = game.Execute(w, &struct {
		Title      string
		Body       []byte
		PlayerName string
	}{
		Title:      "Welcome",
		Body:       body,
		PlayerName: playername,
	})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

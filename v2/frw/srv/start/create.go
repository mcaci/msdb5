package start

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
)

func (s *Session) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	gamename := r.Form["gamename"][0]
	if s.Game.Created(gamename) {
		log.Print("another game already exists, cannot create more than 1")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	s.Game = briscola.NewGame(&briscola.Options{
		WithName: gamename,
	})
	playername := r.Form["playername"][0]
	err = briscola.Register(playername, s.Game)
	if err != nil {
		log.Print("registration error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("player %q joining game %q", playername, gamename)
}

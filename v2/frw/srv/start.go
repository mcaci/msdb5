package srv

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/frw/session"
	"github.com/mcaci/msdb5/v2/frw/srv/assets"
)

var (
	s    = session.NewBriscola()
	s5   = session.Briscola5{}
	game = template.Must(template.ParseFiles("assets/game.html"))
)

func Start(w http.ResponseWriter, r *http.Request) {
	go startOption(w, r)
	s.Wg.Wait()

	if r.FormValue("type") == "create" {
		briscola.Start(s.Game)
		log.Print(s.Game)
		s.Wg.Add(2)
	}

	pl := s.Game.Players().At(s.GetAndIncr())
	assets.MustExecute(assets.Game, w, &struct{ PlayerName interface{} }{PlayerName: pl.Name()})
	assets.MustExecute(assets.Hand(pl), w, nil)
	assets.MustExecute(assets.Label("Briscola"), w, &struct{ Label interface{} }{Label: s.Game.Briscola()})
	assets.MustExecute(assets.Label("Player"), w, &struct{ Label interface{} }{Label: pl})
}

func startOption(w http.ResponseWriter, r *http.Request) {
	defer s.Wg.Done()
	switch r.FormValue("type") {
	case "create":
		s.Create(w, r)
	case "join":
		s.Join(w, r)
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "did not understand the action", http.StatusInternalServerError)
	}
}

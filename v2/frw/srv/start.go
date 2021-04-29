package srv

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	briscolad "github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/frw/session"
)

var (
	s    = session.NewBriscola()
	s5   = session.Briscola5{}
	game = template.Must(template.ParseFiles("assets/game.html"))
)

func Start(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("type") {
	case "create":
		create(w, r)
	case "join":
		join(w, r)
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "did not understand the action", http.StatusInternalServerError)
		return
	}
	plId := s.NPls
	s.NPls++
	switch session.NPlBriscola {
	case int(s.NPls):
		briscola.StartGame(s.Game)
		session.Signal(s.Ready)
		s.NPls = 0
	default:
		session.Wait(s.Ready)
	}
	log.Print(s.Game)
	err := game.Execute(w, data(plId))
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	gamename := r.Form["gamename"][0]
	if s.Game.Started(gamename) {
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

func join(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	gamename := r.Form["gamename"][0]
	if !s.Game.Started(gamename) {
		log.Printf("game %s not found", gamename)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	playername := r.Form["playername"][0]
	err = briscola.Register(playername, s.Game)
	if err != nil {
		log.Print("registration error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("player %q joining game %q", playername, gamename)
}

func data(plId uint8) interface{} {
	pl := s.Game.Players().At(int(plId))
	return &struct {
		Title      string
		Player     string
		Hand       set.Cards
		Briscola   *briscolad.Card
		Board      string
		PlayerName string
		NextPlayer string
	}{
		Title:      "Welcome",
		Player:     pl.String(),
		Hand:       *pl.Hand(),
		Briscola:   s.Game.Briscola(),
		PlayerName: pl.Name(),
	}
}

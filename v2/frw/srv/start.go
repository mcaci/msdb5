package srv

import (
	"html/template"
	"log"
	"net/http"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	briscolad "github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/frw/session"
	"github.com/mcaci/msdb5/v2/frw/srv/start"
)

var (
	s    = session.NewBriscola()
	s5   = session.Briscola5{}
	game = template.Must(template.ParseFiles("assets/game.html"))
)

func Start(w http.ResponseWriter, r *http.Request) {
	sb := start.Session(*s)
	switch r.FormValue("type") {
	case "create":
		sb.Create(w, r)
	case "join":
		sb.Join(w, r)
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "did not understand the action", http.StatusInternalServerError)
		return
	}
	sbTmp := session.Briscola(sb)
	s = &sbTmp
	plId := s.NPls
	s.NPls++
	switch session.NPlBriscola {
	case int(s.NPls):
		briscola.Start(s.Game)
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

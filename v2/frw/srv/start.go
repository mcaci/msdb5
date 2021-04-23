package srv

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/frw/session"
)

var (
	s    = session.Briscola{}
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
	var body string
	var briscolaCard string
	switch r.Form["type"][0] {
	case "create":
		if s.Game != nil && s.Game.Started(gamename) {
			log.Print("another game already exists, cannot create more than 1")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		s.Game = briscola.NewGame(&briscola.Options{
			WithName: gamename,
		})
		err := briscola.Register(playername, s.Game)
		if err != nil {
			log.Print("registration error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		body = fmt.Sprintf("new game created with gamename %q by player %q", gamename, playername)
		log.Print(string(body))
		s.NPls++
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
		body = fmt.Sprintf("player %q joining game %q", playername, gamename)
		log.Print(string(body))
		s.NPls++
		if s.NPls == session.NPlBriscola {
			briscola.StartGame(s.Game)
		}
		briscolaCard = s.Game.Briscola().String()
	default:
		log.Printf("unknown %q option", r.Form["type"][0])
		http.Error(w, "did not understand the action", http.StatusInternalServerError)
		return
	}
	i, err := s.Game.Players().Players.Index(func(p *player.Player) bool { return p.Name() == playername })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pl := s.Game.Players().At(int(i))
	log.Print(s.Game)
	err = game.Execute(w, &struct {
		Title      string
		Body       string
		Hand       set.Cards
		Briscola   string
		Board      string
		PlayerName string
	}{
		Title:      "Welcome",
		Body:       pl.String(),
		Hand:       *pl.Hand(),
		Briscola:   briscolaCard,
		PlayerName: playername,
	})
	if err != nil {
		http.NotFound(w, r)
		return
	}
}

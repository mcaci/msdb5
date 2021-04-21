package srv

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/dom/player"
)

func Play(w http.ResponseWriter, r *http.Request) {
	m := validName.FindStringSubmatch(r.URL.Path)

	playername := m[2]
	cardname := r.FormValue("cardname")
	i, err := s.Game.Players().Players.Index(func(p *player.Player) bool { return p.Name() == playername })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pl := s.Game.Players().At(int(i))
	log.Print("playing card", cardname)
	err = game.Execute(w, &struct {
		Title      string
		Body       string
		Briscola   string
		Board      string
		PlayerName string
	}{
		Title:      "Player",
		Body:       pl.String(),
		Briscola:   s.Game.Briscola().String(),
		Board:      fmt.Sprint(*s.Game.Board()),
		PlayerName: pl.Name(),
	})
	log.Print(s.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

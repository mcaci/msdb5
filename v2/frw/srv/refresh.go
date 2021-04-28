package srv

import (
	"log"
	"net/http"
	"regexp"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
)

var (
	validName = regexp.MustCompile("^/(refresh|play)/([a-zA-Z0-9]+)$")
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	m := validName.FindStringSubmatch(r.URL.Path)
	playername := m[2]
	i, err := s.Game.Players().Players.Index(func(p *player.Player) bool { return p.Name() == playername })
	pl := s.Game.Players().At(int(i))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = game.Execute(w, &struct {
		Title      string
		Player     string
		Hand       set.Cards
		Briscola   *briscola.Card
		Board      interface{}
		PlayerName string
		NextPlayer string
	}{
		Title:      "Player",
		Player:     pl.String(),
		Hand:       *pl.Hand(),
		PlayerName: pl.Name(),
		Briscola:   s.Game.Briscola(),
		Board:      *s.Game.Board().Cards,
		NextPlayer: s.Game.Players().Players[s.Curr].Name(),
	})
	log.Print(s.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
	validName = regexp.MustCompile("^/(draw|play)/([a-zA-Z0-9]+)$")
)

func Draw(w http.ResponseWriter, r *http.Request) {
	m := validName.FindStringSubmatch(r.URL.Path)
	playername := m[2]
	i, err := s.Game.Players().Players.Index(func(p *player.Player) bool { return p.Name() == playername })
	pl := s.Game.Players().At(int(i))
	// pl.Hand().Add(s.Game.Deck().Top())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = game.Execute(w, &struct {
		Title      string
		Body       string
		Hand       set.Cards
		Briscola   *briscola.Card
		Board      string
		PlayerName string
		NextPlayer string
	}{
		Title:      "Player",
		Body:       pl.String(),
		Hand:       *pl.Hand(),
		Briscola:   s.Game.Briscola(),
		PlayerName: pl.Name(),
	})
	log.Print(s.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

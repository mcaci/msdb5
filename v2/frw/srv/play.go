package srv

import (
	"log"
	"net/http"

	"github.com/mcaci/msdb5/v2/dom/player"
)

// var (
// validPlay = regexp.MustCompile("^/play/([1-9]|10)/([a-zA-Z0-9](O|C|S|B))$")
// )

func Play(w http.ResponseWriter, r *http.Request) {
	m := validName.FindStringSubmatch(r.URL.Path)

	playername := m[2]
	cardname := r.FormValue("cardname")
	i, err := s.Game.Players().Players.Index(func(p *player.Player) bool { return p.Name() == playername })
	pl := s.Game.Players().At(int(i))
	log.Print("playing card", cardname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = game.Execute(w, &struct {
		Title      string
		Body       string
		PlayerName string
	}{
		Title:      "Player",
		Body:       pl.String(),
		PlayerName: pl.Name(),
	})
	log.Print(s.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

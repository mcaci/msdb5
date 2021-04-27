package srv

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
)

func Play(w http.ResponseWriter, r *http.Request) {
	m := validName.FindStringSubmatch(r.URL.Path)

	playername := m[2]
	cardn, err := strconv.Atoi(r.FormValue("cardn"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	i, err := s.Game.Players().Players.Index(func(p *player.Player) bool { return p.Name() == playername })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pl := s.Game.Players().At(int(i))
	card, err := pl.Select(cardn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("playing card %v", card)
	board := fmt.Sprint(*s.Game.Board().Cards)
	info := play.Round(&play.RoundOpts{
		PlIdx:        i,
		PlHand:       pl.Hand(),
		CardIdx:      uint8(cardn),
		PlayedCards:  s.Game.Board(),
		NPlayers:     2,
		BriscolaCard: *s.Game.Briscola(),
		EndRound:     play.EndDirect,
	})
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
		Board:      board,
		PlayerName: pl.Name(),
		NextPlayer: s.Game.Players().Players[info.NextPl].Name(),
	})
	log.Print(s.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if !info.NextRnd {
		return
	}
	briscola.Collect(info.OnBoard, s.Game.Players().At(int(info.NextPl)))
}

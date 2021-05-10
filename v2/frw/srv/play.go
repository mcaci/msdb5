package srv

import (
	"log"
	"net/http"
	"strconv"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/frw/session"
	"github.com/mcaci/msdb5/v2/pb"
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
	card, err := pl.SelectCard(cardn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("playing card %v", card)
	info := play.Round(&play.RoundOpts{
		PlIdx:        i,
		PlHand:       pl.Hand(),
		CardIdx:      uint8(cardn),
		PlayedCards:  s.Game.Board(),
		NPlayers:     2,
		BriscolaCard: *s.Game.Briscola(),
		EndRound:     endDirect,
	})
	s.Game.Board().Cards = info.OnBoard.Cards
	s.Curr = info.NextPl
	s.NPls++

	switch session.NPlBriscola {
	case int(s.NPls):
		session.Signal(s.Ready)
		s.NPls = 0
	default:
		session.Wait(s.Ready)
	}

	log.Print(s.Game)
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
		Board:      *info.OnBoard.Cards,
		NextPlayer: s.Game.Players().Players[s.Curr].Name(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pl.Hand().Add(s.Game.Deck().Top())
	if !info.NextRnd {
		return
	}
	briscola.Collect(info.OnBoard, s.Game.Players().At(int(info.NextPl)))
}

func endDirect(opts *struct {
	PlayedCards  briscola.PlayedCards
	BriscolaCard briscola.Card
}) (*pb.Index, error) {
	pbcards := make(set.Cards, len(*opts.PlayedCards.Cards))
	for i := range pbcards {
		pbcards[i] = (*opts.PlayedCards.Cards)[i]
	}
	return &pb.Index{Id: uint32(briscola.Winner(pbcards, opts.BriscolaCard.Seed()))}, nil
}

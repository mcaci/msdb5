package srv

import (
	"log"
	"net/http"
	"strconv"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/frw/srv/assets"
	"github.com/mcaci/msdb5/v2/pb"
)

func Play(w http.ResponseWriter, r *http.Request) {
	m := validName.FindStringSubmatch(r.URL.Path)

	playername := m[2]
	cardn, err := strconv.Atoi(r.FormValue("cardn"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	i, err := s.Game.Players().List().Index(func(p *player.Player) bool { return p.Name() == playername })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pl := s.Game.Players().At(int(i))
	card, err := pl.SelectCard(cardn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("playing card %v", card)
	var info *play.RoundInfo
	go func() {
		info = play.Round(&play.RoundOpts{
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
		s.Wg.Done()
	}()
	s.Wg.Wait()
	log.Print(s.Game)

	assets.MustExecute(assets.Game, w, &struct{ PlayerName interface{} }{PlayerName: pl.Name()})
	assets.MustExecute(assets.List("Hand", pl.Hand), w, nil)
	assets.MustExecute(assets.Label("Briscola"), w, &struct{ Label interface{} }{Label: s.Game.Briscola()})
	assets.MustExecute(assets.Label("Player"), w, &struct{ Label interface{} }{Label: pl})
	assets.MustExecute(assets.List("Board", s.Game.BoardCards), w, nil)

	pl.Hand().Add(s.Game.Deck().Top())
	if !info.NextRnd {
		return
	}
	briscola.Collect(info.OnBoard, s.Game.Players().At(int(info.NextPl)))
	s.Wg.Add(2)
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

// func data(plId uint8) interface{} {
// 	pl := s.Game.Players().At(int(plId))
// 	return &struct {
// 		Title      string
// 		Player     string
// 		Hand       set.Cards
// 		Briscola   *briscola.Card
// 		Board      interface{}
// 		PlayerName string
// 		NextPlayer string
// 	}{
// 		Title:      "Player",
// 		Player:     pl.String(),
// 		Hand:       *pl.Hand(),
// 		PlayerName: pl.Name(),
// 		Briscola:   s.Game.Briscola(),
// 		Board:      *info.OnBoard.Cards,
// 		NextPlayer: s.Game.Players().At(int(s.Curr)).Name(),
// 	}
// }

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
	idx, err := s.Game.Players().List().Index(func(p *player.Player) bool { return p.Name() == playername })
	continueIfNoErr(w, err)
	cardn, err := strconv.Atoi(r.FormValue("cardn"))
	continueIfNoErr(w, err)

	pl := s.Game.Players().At(int(idx))
	var info *play.RoundInfo
	go func() {
		info = play.Round(&play.RoundOpts{
			PlIdx:        idx,
			PlHand:       pl.Hand(),
			CardIdx:      uint8(cardn),
			PlayedCards:  s.Game.Board(),
			NPlayers:     2,
			BriscolaCard: *s.Game.Briscola(),
			EndRound:     endDirect,
		})
		s.Game.Board().Cards = info.OnBoard.Cards
		s.Curr = info.NextPl
		s.Wg.Done()
	}()
	s.Wg.Wait()

	assets.MustExecute(assets.Header, w, &struct{ PlayerName interface{} }{PlayerName: pl.Name()})
	assets.MustExecute(assets.Label("Briscola"), w, &struct{ Label interface{} }{Label: s.Game.Briscola()})
	assets.MustExecute(assets.List("Board", s.Game.BoardCards), w, nil)

	pl.Hand().Add(s.Game.Deck().Top())
	log.Print(s.Game)
	assets.MustExecute(assets.Play, w, &struct{ PlayerName interface{} }{PlayerName: pl.Name()})
	assets.MustExecute(assets.List("Hand", pl.Hand), w, nil)
	if info.NextRnd {
		briscola.Collect(info.OnBoard, s.Game.Players().At(int(info.NextPl)))
		log.Print(s.Game)
		s.Wg.Add(2)
	}
	assets.MustExecute(assets.Label("Player"), w, &struct{ Label interface{} }{Label: pl})
	assets.MustExecute(assets.Label("Next Player"), w, &struct{ Label interface{} }{Label: s.Game.Players().At(int(s.Curr)).Name()})
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

func continueIfNoErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

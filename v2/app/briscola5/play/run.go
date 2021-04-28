package play

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/mcaci/msdb5/v2/app/briscola5/end"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(g struct {
	Players      briscola5.Players
	BriscolaCard briscola.Card
}) struct {
	OnBoard briscola.PlayedCards
} {
	playedCards := briscola.NewPlayedCards(5)
	plIdx, err := currentPlayerIndex(g.Players.Caller(), briscola5.ToGeneralPlayers(g.Players))
	if err != nil {
		log.Fatal("didn't expect to arrive at this point")
	}

	for !end.Cond(&end.Opts{
		PlayedCards:  *playedCards,
		Players:      g.Players,
		BriscolaCard: g.BriscolaCard,
	}) {
		rand.Seed(time.Now().Unix())
		hnd := g.Players.At(int(plIdx)).Hand()
		info := Round(&RoundOpts{
			PlHand:       hnd,
			PlIdx:        plIdx,
			CardIdx:      uint8(rand.Intn(len(*hnd))),
			PlayedCards:  playedCards,
			NPlayers:     uint8(len(briscola5.ToGeneralPlayers(g.Players))),
			BriscolaCard: g.BriscolaCard,
			EndRound:     EndRemote,
		})
		playedCards = info.OnBoard
		plIdx = info.NextPl
		if !info.NextRnd {
			continue
		}
		briscola.Collect(playedCards, g.Players.At(int(plIdx)))
	}
	return struct{ OnBoard briscola.PlayedCards }{
		OnBoard: *playedCards,
	}
}

func currentPlayerIndex(cp *player.Player, pls team.Players) (uint8, error) {
	for i := range pls {
		if pls[i] != cp {
			continue
		}
		return uint8(i), nil
	}
	return 0, errors.New("not found")
}

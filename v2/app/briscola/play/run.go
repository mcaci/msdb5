package play

import (
	"log"
	"math/rand"
	"time"

	briscolapp "github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/briscola/end"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
)

func Run(g struct {
	Players      misc.Players
	BriscolaCard briscola.Card
	Deck         briscolapp.Deck
	EndRound     func(*struct {
		PlayedCards  briscola.PlayedCards
		BriscolaCard briscola.Card
	}) (*pb.Index, error)
}) struct {
	OnBoard briscola.PlayedCards
} {
	playedCards := briscola.NewPlayedCards(2)
	plIdx, err := g.Players.Index(func(p misc.Player) bool { return p == g.Players[0] })
	if err != nil {
		log.Fatal("didn't expect to arrive at this point")
	}

	for !end.Cond(&end.Opts{Players: g.Players}) {
		rand.Seed(time.Now().Unix())
		info1 := Round(&RoundOpts{
			PlHand:       g.Players[plIdx].Hand(),
			PlIdx:        plIdx,
			CardIdx:      uint8(rand.Intn(len(*g.Players[plIdx].Hand()))),
			PlayedCards:  playedCards,
			NPlayers:     uint8(len(g.Players)),
			BriscolaCard: g.BriscolaCard,
			EndRound:     g.EndRound,
		})
		info2 := Round(&RoundOpts{
			PlHand:       g.Players[info1.NextPl].Hand(),
			PlIdx:        info1.NextPl,
			CardIdx:      uint8(rand.Intn(len(*g.Players[info1.NextPl].Hand()))),
			PlayedCards:  playedCards,
			NPlayers:     uint8(len(g.Players)),
			BriscolaCard: g.BriscolaCard,
			EndRound:     g.EndRound,
		})

		rWin := g.Players[info2.NextPl]
		rLos := g.Players[(info2.NextPl+1)%2]
		briscola.Collect(playedCards, rWin)
		switch len(g.Deck.Cards) {
		case 0:
		case 1:
			rWin.Hand().Add(g.Deck.Top())
			rLos.Hand().Add(g.BriscolaCard.Item)
		default:
			rWin.Hand().Add(g.Deck.Top())
			rLos.Hand().Add(g.Deck.Top())
		}
	}
	return struct{ OnBoard briscola.PlayedCards }{
		OnBoard: *playedCards,
	}
}

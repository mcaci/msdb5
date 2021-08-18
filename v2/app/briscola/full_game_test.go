package briscola

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
)

func TestAiGame2P(t *testing.T) {
	// setup ai game
	g := NewGame(&Options{WithName: "test"})

	pls := g.Players()
	for i := range *pls {
		(*pls)[i] = misc.New(&misc.Options{Name: "Player" + strconv.Itoa(i), For2P: true})
		(*pls)[i].Hand().Add(g.Deck().Top())
		(*pls)[i].Hand().Add(g.Deck().Top())
		(*pls)[i].Hand().Add(g.Deck().Top())
	}
	Set(briscola.Card{Item: g.Deck().Top()}, g)

	// run ai game
	run(struct {
		Players      misc.Players
		BriscolaCard briscola.Card
		Deck         Deck
		EndRound     func(*struct {
			PlayedCards  briscola.PlayedCards
			BriscolaCard briscola.Card
		}) (*pb.Index, error)
	}{
		Players:      *g.Players(),
		BriscolaCard: *g.Briscola(),
		Deck:         *g.Deck(),
		EndRound:     endDirect,
	})

	scoreIn := &struct {
		Players *misc.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: g.Players(),
		Method: func(i int) (interface{ GetPoints() uint32 }, error) {
			p := briscola.Score(*(*g.Players())[i].Pile())
			return p, nil
		},
	}
	log.Println("Score", PrintScore(scoreIn))
}

func run(g struct {
	Players      misc.Players
	BriscolaCard briscola.Card
	Deck         Deck
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

	for !End(&Opts{Players: g.Players}) {
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

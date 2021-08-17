package play

import (
	"log"
	"strconv"
	"testing"

	"github.com/mcaci/ita-cards/set"
	briscolapp "github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
)

func TestAiGame2P(t *testing.T) {
	// setup ai game
	g := briscolapp.NewGame(&briscolapp.Options{WithName: "test"})

	pls := g.Players()
	for i := range *pls {
		(*pls)[i] = misc.New(&misc.Options{Name: "Player" + strconv.Itoa(i), For2P: true})
		(*pls)[i].Hand().Add(g.Deck().Top())
		(*pls)[i].Hand().Add(g.Deck().Top())
		(*pls)[i].Hand().Add(g.Deck().Top())
	}
	briscolapp.Set(briscola.Card{Item: g.Deck().Top()}, g)

	// run ai game
	Run(struct {
		Players      misc.Players
		BriscolaCard briscola.Card
		Deck         briscolapp.Deck
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
	log.Println("Score", briscolapp.PrintScore(scoreIn))
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

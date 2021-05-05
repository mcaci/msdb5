package ai

import (
	"log"
	"testing"

	"github.com/mcaci/ita-cards/set"
	briscolapp "github.com/mcaci/msdb5/v2/app/briscola"
	"github.com/mcaci/msdb5/v2/app/briscola5"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	briscola5d "github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/pb"
)

func TestAiGameWithSide(t *testing.T) {
	// setup ai game
	g := briscola5.NewGame(&briscola5.Options{
		WithSide:     true,
		WithCmpF:     dirCmp,
		WithEndRound: endDirect,
	})
	briscola5.SetScoreF(func(i int) (interface{ GetPoints() uint32 }, error) {
		p := briscola.Score(*g.Players().At(i).Pile())
		return p, nil
	}, g)

	// run ai game
	Run(g)

	scoreIn := &struct {
		Players *briscola.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: &briscola.Players{Players: briscola5d.ToGeneralPlayers(*g.Players())},
		Method:  g.ScoreF(),
	}
	log.Println("Score", briscolapp.PrintScore(scoreIn))
}

func TestAiGameWithNoSide(t *testing.T) {
	// setup ai game
	g := briscola5.NewGame(&briscola5.Options{
		WithCmpF:     dirCmp,
		WithEndRound: endDirect,
	})
	briscola5.SetScoreF(func(i int) (interface{ GetPoints() uint32 }, error) {
		p := briscola.Score(*g.Players().At(i).Pile())
		return p, nil
	}, g)

	// run ai game
	Run(g)

	scoreIn := &struct {
		Players *briscola.Players
		Method  func(int) (interface{ GetPoints() uint32 }, error)
	}{
		Players: &briscola.Players{Players: briscola5d.ToGeneralPlayers(*g.Players())},
		Method:  g.ScoreF(),
	}
	log.Println("Score", briscolapp.PrintScore(scoreIn))
}

func dirCmp(curr, prop briscola5d.AuctionScore) int8 {
	return int8(briscola5d.Cmp(curr, prop))
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

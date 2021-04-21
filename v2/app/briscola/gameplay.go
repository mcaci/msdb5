package briscola

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/app/briscola/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func Start(g *Game) {
	StartGame(g)

	// play phase
	play.Run(struct {
		Players      briscola.Players
		BriscolaCard briscola.Card
	}{
		Players:      g.players,
		BriscolaCard: g.briscolaCard,
	})
}

func StartGame(g *Game) {
	// distribute cards to players
	distributeCards(g)

	// set briscola card
	g.briscolaCard = briscola.Card{Item: g.deck.Top()}
}

func Score(g *Game) string {
	return fmt.Sprintf("[%s: %d], [%s: %d]",
		"Player 1", briscola.Score(*g.players.Players[0].Pile()),
		"Player 2", briscola.Score(*g.players.Players[1].Pile()))
}

const NPlBriscola = 2

func distributeCards(g *Game) {
	for i := 0; i < 3; i++ {
		g.players.At(0).Hand().Add(g.deck.Top())
		g.players.At(1).Hand().Add(g.deck.Top())
	}
}

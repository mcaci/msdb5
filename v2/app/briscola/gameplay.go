package briscola

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/briscola/play"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func Start(g *Game) {
	// distribute cards to players
	distributeCards(g)

	g.briscolaCard = briscola.Card{Item: *card.MustID(1)}

	// play phase
	play.Run(struct {
		Players      briscola.Players
		BriscolaCard briscola.Card
	}{
		Players:      g.players,
		BriscolaCard: g.briscolaCard,
	})
}

func Score(g *Game) string {
	return fmt.Sprintf("[%s: %d], [%s: %d]",
		"Player 1", briscola.Score(*g.players.Players[0].Pile()),
		"Player 2", briscola.Score(*g.players.Players[1].Pile()))
}

func distributeCards(g *Game) {
	d := set.Deck()
	for i := 0; i < set.DeckSize; i++ {
		g.players.At(i % 2).Hand().Add(d.Top())
	}
}

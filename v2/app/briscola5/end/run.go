package end

import (
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func Run(g struct {
	PlayedCards  briscola.PlayedCards
	Players      misc.Players
	BriscolaCard briscola.Card
	Side         briscola5.Side
}) {
	// give all left cards to the player with highest card value for briscola
	var nextPlayer uint8
	for _, card := range briscola.Serie(g.BriscolaCard) {
		i, err := g.Players.Index(misc.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		nextPlayer = i
	}

	// collect cards
	briscola.Collect(newAllCards(g.Players, g.Side, &g.PlayedCards), g.Players[nextPlayer])
}

package end

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(g struct {
	PlayedCards  set.Cards
	Players      team.Players
	BriscolaCard interface{ Seed() card.Seed }
	Side         set.Cards
}) {
	// no more cards to play
	if g.Players.All(player.EmptyHanded) {
		return
	}

	// give all left cards to the player with highest card value for briscola
	var nextPlayer uint8
	for _, card := range serie(g.BriscolaCard.Seed()) {
		i, err := g.Players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		nextPlayer = i
	}

	// collect cards
	set.Move(collect.NewAllCards(g.Players, &g.Side, &g.PlayedCards).Set(), g.Players[nextPlayer].Pile())
}

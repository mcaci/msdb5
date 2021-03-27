package game

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runEnd(g struct {
	players      team.Players
	briscolaCard card.Item
	playedCards  set.Cards
	side         set.Cards
}) {
	// no more cards to play
	if g.players.All(player.EmptyHanded) {
		return
	}

	// give all left cards to the player with highest card value for briscola
	var nextPlayer uint8
	for _, card := range serie(g.briscolaCard.Seed()) {
		i, err := g.players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		nextPlayer = i
	}

	// collect cards
	set.Move(collect.NewAllCards(g.players, &g.side, &g.playedCards).Set(), g.players[nextPlayer].Pile())
}

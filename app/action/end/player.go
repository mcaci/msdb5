package end

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type currentPlayerProvider interface {
	CurrentPlayer() *player.Player
	Players() team.Players
	Briscola() card.Item
}

func LastPlayer(g currentPlayerProvider) *player.Player {
	lastPl := g.CurrentPlayer()
	for _, card := range briscola.Serie(g.Briscola()) {
		_, p := g.Players().Find(player.IsCardInHand(card))
		if p == nil { // no one has card
			continue
		}
		lastPl = p
		break
	}
	return lastPl
}

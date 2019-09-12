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

// LastPlayer func
func LastPlayer(g currentPlayerProvider) *player.Player {
	for _, c := range briscola.Serie(g.Briscola().Seed()) {
		if _, p := g.Players().Find(player.IsCardInHand(c)); p != nil {
			return p
		}
	}
	return g.CurrentPlayer()
}

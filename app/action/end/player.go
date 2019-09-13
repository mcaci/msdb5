package end

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type currentPlayerCardsProvider interface {
	CurrentPlayer() *player.Player
	Cards() *set.Cards
}

// LastPlayer func
func LastPlayer(g currentPlayerCardsProvider, players team.Players) *player.Player {
	for _, c := range *g.Cards() {
		if _, p := players.Find(player.IsCardInHand(c)); p != nil {
			return p
		}
	}
	return g.CurrentPlayer()
}

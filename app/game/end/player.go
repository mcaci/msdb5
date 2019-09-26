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
func LastPlayer(cards set.Cards, players team.Players, curr *player.Player) *player.Player {
	for _, c := range cards {
		if _, p := players.Find(player.IsCardInHand(c)); p != nil {
			return p
		}
	}
	return curr
}

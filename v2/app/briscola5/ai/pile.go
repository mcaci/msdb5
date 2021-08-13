package ai

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/misc"
)

// CommonPile returns the pile of cards of all players in input
func CommonPile(players misc.Players) set.Cards {
	pile := make(set.Cards, 0)
	for _, p := range players {
		pile.Add(*p.Pile()...)
	}
	return pile
}

package briscola

import "github.com/mcaci/ita-cards/set"

// CommonPile returns the pile of cards of all players in input
func CommonPile(players Players) set.Cards {
	pile := make(set.Cards, 0)
	for _, p := range players {
		pile.Add(*p.Pile()...)
	}
	return pile
}

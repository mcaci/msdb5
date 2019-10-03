package team

import "github.com/mcaci/ita-cards/set"

// CommonPile func
func CommonPile(players Players) set.Cards {
	pile := make(set.Cards, 0)
	for _, p := range players {
		pile.Add(*p.Pile()...)
	}
	return pile
}

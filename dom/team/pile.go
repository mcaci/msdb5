package team

import "github.com/mcaci/ita-cards/set"

// CommonPile func
func CommonPile(pls Players) set.Cards {
	pile := make(set.Cards, 0)
	for _, p := range pls {
		pile.Add(*p.Pile()...)
	}
	return pile
}

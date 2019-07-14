package briscola

import (
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
)

type seeder interface {
	Seed() card.Seed
}

// Serie func
func Serie(briscola seeder) deck.Cards {
	set := deck.Cards{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	if briscola.Seed() != card.Coin {
		for i := range set {
			set[i] += card.ID(10 * briscola.Seed())
		}
	}
	return set
}

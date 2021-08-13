package briscola

import (
	"github.com/mcaci/ita-cards/set"
)

type Deck struct{ set.Cards }

func NewDeck() *Deck {
	return &Deck{Cards: set.Deck()}
}

func Distribute(g *struct {
	Players  misc.Players
	Deck     *Deck
	HandSize int
}) {
	for i := 0; i < g.HandSize; i++ {
		for _, p := range g.Players {
			p.Hand().Add(g.Deck.Top())
		}
	}
}

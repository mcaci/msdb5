package player

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Player interface
type Player interface {
	Draw(deck.Deck) card.Card
	Name() string
	Hand() card.Cards
	fmt.Stringer

	Iam(string)
	MyHostIs(string)

	Has(c card.Card) bool
}

// New func
func New() Player {
	player := new(concretePlayer)
	player.hand = card.Cards{}
	return player
}

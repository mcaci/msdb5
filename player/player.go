package player

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// Player interface
type Player interface {
	Draw(func() card.Card) card.Card
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

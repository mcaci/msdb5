package player

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// Player interface
type Player interface {
	Draw(card.Supplier) uint8
	Name() string
	Hand() *card.Cards
	fmt.Stringer

	Iam(string)
	MyHostIs(string)

	Has(uint8) bool
}

// New func
func New() Player {
	player := new(concretePlayer)
	player.hand = card.Cards{}
	return player
}

package player

import (
	"container/list"
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Player interface
type Player interface {
	Draw(deck.Deck) *card.Card
	Play() *card.Card
	Name() string
	Hand() *list.List
	fmt.Stringer

	Iam(string)
	MyHostIs(string)

	has(c *card.Card) bool
}

// New func
func New() Player {
	player := new(concretePlayer)
	player.hand = new(list.List)
	return player
}

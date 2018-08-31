package player

import (
	"container/list"
	"msdb5/card"
	"msdb5/deck"
)

// Player interface
type Player interface {
	Draw(d deck.Deck) *card.Card
	Has(c *card.Card) bool
	Hasnt(c *card.Card) bool
}

// New func
func New() Player {
	player := new(concretePlayer)
	player.cards = new(list.List)
	return player
}

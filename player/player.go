package player

import (
	"container/list"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Player interface
type Player interface {
	Draw(d deck.Deck) *card.Card
	Play() *card.Card
	
	has(c *card.Card) bool
}

// New func
func New() Player {
	player := new(concretePlayer)
	player.cards = new(list.List)
	return player
}

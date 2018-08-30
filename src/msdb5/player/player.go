package player

import "container/list"
import "msdb5/deck"
import "msdb5/card"

// Player interface
type Player interface {
	Init()
	Draw(d deck.Deck) *card.Card
	Has(c *card.Card) bool
	Hasnt(c *card.Card) bool
}

// ConcretePlayer type
type ConcretePlayer struct {
	cards *list.List
}

// Init func
func (player *ConcretePlayer) Init() {
	player.cards = list.New()
}

// Draw func
func (player *ConcretePlayer) Draw(d deck.Deck) *card.Card {
	c := d.RemoveTop()
	player.cards.PushFront(c)
	return c
}

// Has func
func (player *ConcretePlayer) Has(c *card.Card) bool {
	cardFound := false
	for e := player.cards.Front(); e != nil; e = e.Next() {
		cardFound = (e.Value == c)
	}
	return cardFound
}

// Hasnt func
func (player *ConcretePlayer) Hasnt(c *card.Card) bool {
	return !player.Has(c)
}

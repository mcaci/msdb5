package msdb5

import "container/list"

type Player interface {
	New()
	Draw(d Deck) *Card
	Has(c *Card) bool
	Hasnt(c *Card) bool
}

type ConcretePlayer struct {
	cards *list.List
}

func (player *ConcretePlayer) New() {
	player.cards = list.New()
}

func (player *ConcretePlayer) Draw(d Deck) *Card {
	c := d.First()
	player.cards.PushFront(c)
	return c
}

func (player *ConcretePlayer) Has(c *Card) bool {
	cardFound := false
	for e := player.cards.Front(); e != nil; e = e.Next() {
		cardFound = (e.Value == c)
	}
	return cardFound
}

func (player *ConcretePlayer) Hasnt(c *Card) bool {
	return !player.Has(c)
}

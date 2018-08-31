package player

import "container/list"
import "msdb5/deck"
import "msdb5/card"

type concretePlayer struct {
	cards *list.List
}

// Draw func
func (player *concretePlayer) Draw(d deck.Deck) *card.Card {
	c := d.RemoveTop()
	player.cards.PushFront(c)
	return c
}

// Has func
func (player *concretePlayer) Has(c *card.Card) bool {
	cardFound := false
	for e := player.cards.Front(); e != nil; e = e.Next() {
		cardFound = (e.Value == c)
	}
	return cardFound
}

// Hasnt func
func (player *concretePlayer) Hasnt(c *card.Card) bool {
	return !player.Has(c)
}

func (player concretePlayer) String() string {
	str := "concretePlayer["
	for e := player.cards.Front(); e != nil; e = e.Next() {
		str += e.Value.(*card.Card).String() + " "
	}
	str += "]"
	return str
}

package player

import "container/list"
import "github.com/nikiforosFreespirit/msdb5/deck"
import "github.com/nikiforosFreespirit/msdb5/card"

type concretePlayer struct {
	cards *list.List
}

// Draw func
func (player *concretePlayer) Draw(d deck.Deck) *card.Card {
	c := d.RemoveTop()
	player.cards.PushFront(c)
	return c
}

// Play func
func (player *concretePlayer) Play() *card.Card {
	return player.cards.Front().Value.(*card.Card)
}

func (player *concretePlayer) Name() string {
	return "Michele"
}

func (player *concretePlayer) has(c *card.Card) bool {
	cardFound := false
	for e := player.cards.Front(); e != nil; e = e.Next() {
		cardFound = (e.Value == c)
	}
	return cardFound
}

func (player concretePlayer) String() string {
	str := "concretePlayer["
	for e := player.cards.Front(); e != nil; e = e.Next() {
		str += e.Value.(*card.Card).String() + " "
	}
	str += "]"
	return str
}

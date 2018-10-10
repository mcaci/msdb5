package player

import "container/list"
import "github.com/nikiforosFreespirit/msdb5/deck"
import "github.com/nikiforosFreespirit/msdb5/card"

type concretePlayer struct {
	name string
	host string
	hand *list.List
}

// Draw func
func (player *concretePlayer) Draw(d deck.Deck) card.Card {
	c := d.RemoveTop()
	player.hand.PushFront(c)
	return c
}

func (player *concretePlayer) Hand() *list.List {
	return player.hand
}

func (player *concretePlayer) Name() string {
	return player.name
}

func (player *concretePlayer) Iam(name string) {
	player.name = name
}

func (player *concretePlayer) MyHostIs(host string) {
	player.host = host
}

func (player *concretePlayer) Has(c card.Card) bool {
	var cardFound bool
	for e := player.hand.Front(); e != nil; e = e.Next() {
		cardFound = (e.Value == c)
	}
	return cardFound
}

func (player concretePlayer) String() string {
	str := "concretePlayer["
	str += "Name:" + player.name + ";"
	str += "Host:" + player.host + ";"
	for e := player.hand.Front(); e != nil; e = e.Next() {
		str += e.Value.(*card.Card).String() + " "
	}
	str += "]"
	return str
}

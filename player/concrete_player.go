package player

import "github.com/nikiforosFreespirit/msdb5/deck"
import "github.com/nikiforosFreespirit/msdb5/card"

type concretePlayer struct {
	name string
	host string
	hand []card.Card
}

// Draw func
func (player *concretePlayer) Draw(d deck.Deck) card.Card {
	c := d.RemoveTop()
	player.hand = append(player.hand, c)
	return c
}

func (player *concretePlayer) Hand() []card.Card {
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
	for _, card := range player.hand {
		cardFound = (c == card)
		if cardFound {
			break
		}
	}
	return cardFound
}

func (player concretePlayer) String() string {
	str := "concretePlayer["
	str += "Name:" + player.name + ";"
	str += "Host:" + player.host + ";"
	for _, card := range player.hand {
		str += card.String() + " "
	}
	str += "]"
	return str
}

package player

import "github.com/nikiforosFreespirit/msdb5/card"

type concretePlayer struct {
	name string
	host string
	hand card.Cards
}

// Draw func
func (player *concretePlayer) Draw(cardSupplier card.Supplier) card.Card {
	c := cardSupplier.Supply()
	player.Hand().Add(c)
	return c
}

func (player *concretePlayer) Hand() *card.Cards {
	return &player.hand
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
	return player.Hand().Has(c)
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

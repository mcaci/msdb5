package player

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

type concretePlayer struct {
	name string
	host string
	hand card.Cards
}

// Draw func
func (player *concretePlayer) Draw(cardSupplier card.Supplier) uint8 {
	c := cardSupplier.Supply()
	player.Hand().AddID(c)
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
	return player.HasID(c.ID())
}

func (player *concretePlayer) HasID(id uint8) bool {
	return player.Hand().HasID(id)
}

func (player concretePlayer) String() string {
	str := "concretePlayer["
	str += "Name:" + player.name + ";"
	str += "Host:" + player.host + ";"
	for _, cardID := range player.hand {
		c, _ := card.ByID(cardID)
		str += c.String() + " "
	}
	str += "]"
	return str
}

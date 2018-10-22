package player

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Player struct
type Player struct {
	name string
	host string
	hand card.Cards
}

// New func
func New() *Player {
	player := new(Player)
	player.hand = card.Cards{}
	return player
}

// Draw func
func (player *Player) Draw(cardSupplier card.Supplier) card.ID {
	c := cardSupplier.Supply()
	player.Hand().Add(c)
	return c
}

// Hand func
func (player *Player) Hand() *card.Cards {
	return &player.hand
}

// Name func
func (player *Player) Name() string {
	return player.name
}

// Iam func
func (player *Player) Iam(name string) {
	player.name = name
}

// MyHostIs func
func (player *Player) MyHostIs(host string) {
	player.host = host
}

// Has func
func (player *Player) Has(id card.ID) bool {
	return player.Hand().Has(id)
}

func (player Player) String() string {
	str := "Player["
	str += "Name:" + player.name + ";"
	str += "Host:" + player.host + ";"
	for _, cardID := range player.hand {
		c, _ := card.ByID(cardID)
		str += c.String() + " "
	}
	str += "]"
	return str
}

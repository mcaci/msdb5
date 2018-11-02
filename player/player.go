package player

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Player struct
type Player struct {
	name string
	host string
	hand card.Cards
	pile card.Cards
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

// SetName func
func (player *Player) SetName(name string) {
	player.name = name
}

// MyHostIs func
func (player *Player) MyHostIs(host string) {
	player.host = host
}

// Host func
func (player *Player) Host() string {
	return player.host
}

// Has func
func (player *Player) Has(id card.ID) bool {
	return player.Hand().Has(id)
}

// Fold func
func (player *Player) Fold() bool {
	return true
}

// Pile func
func (player *Player) Pile() *card.Cards {
	return &player.pile
}

// Collect func
func (player *Player) Collect(cards card.Cards) {
	player.pile.Add(cards...)
}

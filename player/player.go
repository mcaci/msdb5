package player

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
)

// Player struct
type Player struct {
	name string
	host string
	hand set.Cards
	pile set.Cards
}

// New func
func New() *Player {
	player := new(Player)
	player.hand = set.Cards{}
	return player
}

// Draw func
func (player *Player) Draw(cardSupplier card.Supplier) card.ID {
	c := cardSupplier.Supply()
	player.Hand().Add(c)
	return c
}

// Hand func
func (player *Player) Hand() *set.Cards {
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
func (player *Player) Pile() *set.Cards {
	return &player.pile
}

// Collect func
func (player *Player) Collect(cards set.Cards) {
	if len(cards) > 0 {
		player.Pile().Add(cards...)
	}
}

// Score func
func (player *Player) Score(count func(cards set.Cards) uint8) uint8 {
	return count(*player.Pile())
}

// Supply func
func (player *Player) Supply() card.ID {
	return 1
}

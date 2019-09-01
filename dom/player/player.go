package player

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
)

// Player struct
type Player struct {
	name, host string
	hand, pile set.Cards
	fold       bool
	info       chan []byte
}

// New func
func New() *Player {
	return new(Player)
}

// Name func
func (player Player) Name() string { return player.name }

// Hand func
func (player *Player) Hand() *set.Cards { return &player.hand }

// Pile func
func (player *Player) Pile() *set.Cards { return &player.pile }

// RegisterAs func
func (player *Player) RegisterAs(name string) { player.name = name }

// Join func
func (player *Player) Join(origin string) { player.host = origin }

// Attach func
func (player *Player) Attach(info chan []byte) { player.info = info }

// Fold func
func (player *Player) Fold() { player.fold = true }

// Write func
func (player *Player) Write(msg []byte) (n int, err error) {
	player.info <- []byte(msg)
	return len(msg), nil
}

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		player.name, player.hand, player.pile, player.fold)
}

package player

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
)

// B2Player struct
type B2Player struct {
	name       string
	hand, pile set.Cards
}

// B5Player struct
type B5Player struct {
	B2Player
	fold bool
}

// Options struct
type Options struct {
	Name  string
	For2P bool
	For5P bool
}

// Player interface
type Player interface {
	Name() string
	Hand() *set.Cards
	Pile() *set.Cards
}

// New func
func New(o *Options) Player {
	b2P := B2Player{name: o.Name}
	var p Player
	switch {
	case o.For2P:
		p = &b2P
	case o.For5P:
		p = &B5Player{B2Player: b2P}
	}
	return p
}

// Name func
func (player B2Player) Name() string { return player.name }

// Hand func
func (player *B2Player) Hand() *set.Cards { return &player.hand }

// Pile func
func (player *B2Player) Pile() *set.Cards { return &player.pile }

// RegisterAs func
func (player *B2Player) RegisterAs(name string) { player.name = name }

func (player B2Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v)",
		player.name, player.hand, player.pile)
}

// Fold func
func (player *B5Player) Fold() { player.fold = true }

func (player B5Player) String() string {
	return fmt.Sprintf("(Player: %+v, Has folded? %t)\n", player.B2Player, player.fold)
}

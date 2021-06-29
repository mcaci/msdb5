package player

import (
	"fmt"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

// Player struct
type Player struct {
	name       string
	hand, pile set.Cards
}

// Options struct
type Options struct {
	For2P bool
	For5P bool
}

// Playerer interface
type Playerer interface{}

// NewWithOpts func
func NewWithOpts(o *Options) Playerer {
	switch {
	case o.For2P:
		return &Player{}
	case o.For5P:
		return &B5Player{}
	default:
		return &Player{}
	}
}

// New func
func New() *Player { return &Player{} }

// Name func
func (player Player) Name() string { return player.name }

// Hand func
func (player *Player) Hand() *set.Cards { return &player.hand }

// SelectCard func
func (player *Player) SelectCard(i int) (*card.Item, error) {
	if l := len(player.hand); i >= l {
		return nil, fmt.Errorf("card at position %d cannot be found. Maximum position is %d", i, l)
	}
	return &player.hand[i], nil
}

// Pile func
func (player *Player) Pile() *set.Cards { return &player.pile }

// RegisterAs func
func (player *Player) RegisterAs(name string) { player.name = name }

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v)",
		player.name, player.hand, player.pile)
}

// B5Player struct
type B5Player struct {
	Player
	fold bool
}

// NewPlayer func
func NewPlayer() *B5Player { return &B5Player{} }

// Fold func
func (player *B5Player) Fold() { player.fold = true }

func (player B5Player) String() string {
	return fmt.Sprintf("(Player: %+v, Has folded? %t)\n", player.Player, player.fold)
}

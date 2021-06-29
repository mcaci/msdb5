package player

import (
	"fmt"
)

// B5Player struct
type B5Player struct {
	Player
	fold bool
}

// NewPlayer func
func NewPlayer() *B5Player { return &B5Player{} }

// Fold func
func (player *B5Player) Fold() { player.fold = true }

// B5Predicate type
type B5Predicate func(p *B5Player) bool

// Folded var
var Folded B5Predicate = func(p *B5Player) bool { return p.fold }

func (player B5Player) String() string {
	return fmt.Sprintf("(Player: %+v, Has folded? %t)\n", player.Player, player.fold)
}

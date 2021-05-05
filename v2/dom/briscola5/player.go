package briscola5

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/dom/player"
)

// Player struct
type Player struct {
	player.Player
	fold bool
}

// NewPlayer func
func NewPlayer() *Player { return &Player{} }

// Fold func
func (player *Player) Fold() { player.fold = true }

// Predicate type
type Predicate func(p *Player) bool

// Folded var
var Folded Predicate = func(p *Player) bool { return p.fold }

func (player Player) String() string {
	return fmt.Sprintf("(Player: %+v, Has folded? %t)\n", player.Player, player.fold)
}

package briscola5

import (
	"fmt"

	"github.com/mcaci/msdb5/v2/dom/briscola"
)

// Player struct
type Player struct {
	briscola.Player
	fold bool
}

// NewB5Player func
func NewB5Player(name string) *Player { return &Player{Player: *briscola.NewB2Player(name)} }

// Fold func
func (player *Player) Fold() { misc.fold = true }

// Folded func
func (player *Player) Folded() bool { return misc.fold }

func (player Player) String() string {
	return fmt.Sprintf("(Player: %+v, Has folded? %t)\n", misc.Player, misc.fold)
}

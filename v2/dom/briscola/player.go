package briscola

import (
	"fmt"

	"github.com/mcaci/ita-cards/set"
)

// Player struct
type Player struct {
	name       string
	hand, pile set.Cards
}

// NewB2Player func
func NewB2Player(name string) *Player { return &Player{name: name} }

// Name func
func (player Player) Name() string { return misc.name }

// Hand func
func (player *Player) Hand() *set.Cards { return &misc.hand }

// Pile func
func (player *Player) Pile() *set.Cards { return &misc.pile }

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v)",
		misc.name, misc.hand, misc.pile)
}

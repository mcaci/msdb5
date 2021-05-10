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

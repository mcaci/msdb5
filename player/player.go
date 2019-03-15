package player

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Player struct
type Player struct {
	name, host string
	hand       deck.Cards
	pile       deck.Cards
	fold       bool
}

// New func
func New() *Player {
	player := new(Player)
	player.hand = deck.Cards{}
	return player
}

// Draw func
func (player *Player) Draw(cards deck.Cards) {
	player.hand.Add(cards.Supply())
}

// Has func
func (player *Player) Has(id card.ID) bool {
	return player.hand.Has(id)
}

// Hand func
func (player *Player) Hand() *deck.Cards {
	return &player.hand
}

// Join func
func (player *Player) Join(name, origin string) {
	player.name = name
	player.host = origin
}

// Folded func
func (player Player) Folded() bool { return player.fold }

// NotFolded func
func (player Player) NotFolded() bool { return !player.fold }

// IsSame func
func (player *Player) IsSame(other *Player) bool { return player == other }

// IsSameHost func
func (player Player) IsSameHost(origin string) bool { return player.host == origin }

// Name func
func (player Player) Name() string { return player.name }

// IsName func
func (player Player) IsName(name string) bool { return player.name == name }

// IsNameEmpty func
func (player Player) IsNameEmpty() bool { return player.IsName("") }

// IsHandEmpty func
func (player Player) IsHandEmpty() bool { return len(*player.Hand()) == 0 }

// Fold func
func (player *Player) Fold() { player.fold = true }

// Play function
func (player *Player) Play(card card.ID) (err error) {
	index, err := player.Hand().Find(card)
	if err == nil {
		player.Hand().Remove(index)
	}
	return
}

// Collect func
func (player *Player) Collect(cards *deck.Cards) {
	player.pile.Add(*cards...)
}

// Count func
func (player Player) Count(scorer func(card.ID) uint8) uint8 {
	return player.pile.Sum(scorer)
}

// IsExpectedPlayer func
func (player *Player) IsExpectedPlayer(other *Player, origin string) bool {
	return player.IsSame(other) && player.IsSameHost(origin)
}

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %v)",
		player.name, player.hand, player.pile, player.fold)
}

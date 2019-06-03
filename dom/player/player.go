package player

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

// Player struct
type Player struct {
	name, host string
	hand       deck.Cards
	pile       deck.Cards
	fold       bool
	info       chan []byte
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
	_, err := player.hand.Find(id)
	return err == nil
}

// Hand func
func (player *Player) Hand() *deck.Cards {
	return &player.hand
}

// RegisterAs func
func (player *Player) RegisterAs(name string) {
	player.name = name
}

// Join func
func (player *Player) Join(origin string) {
	player.host = origin
}

// Attach func
func (player *Player) Attach(info chan []byte) {
	player.info = info
}

// ReplyWith func
func (player *Player) ReplyWith(message string) {
	player.info <- []byte(message)
}

// Folded func
func (player Player) Folded() bool { return player.fold }

// IsSameHost func
func (player Player) IsSameHost(origin string) bool { return player.host == origin }

// Host func
func (player Player) Host() string { return player.host }

// Name func
func (player Player) Name() string { return player.name }

// IsNameEmpty func
func (player Player) IsNameEmpty() bool { return player.name == "" }

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
	return player.isSame(other) && player.IsSameHost(origin)
}

func (player *Player) isSame(other *Player) bool { return player == other }

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %v)",
		player.name, player.hand, player.pile, player.fold)
}

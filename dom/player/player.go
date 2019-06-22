package player

import (
	"errors"
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
func (player *Player) Draw(supplier func() card.ID) {
	player.hand.Add(supplier())
}

// DropCards func
func (player *Player) DropCards() {
	player.hand.Clear()
}

// Hand func
func (player *Player) Hand() *deck.Cards {
	return &player.hand
}

// HandSize func
func (player *Player) HandSize() int {
	return len(player.hand)
}

// Has func
func (player *Player) Has(id card.ID) bool {
	return player.hand.Find(id) != -1
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

// Write func
func (player *Player) Write(msg []byte) (n int, err error) {
	player.info <- []byte(msg)
	return len(msg), nil
}

// Folded func
func (player Player) Folded() bool { return player.fold }

// IsSameHost func
func (player Player) IsSameHost(origin string) bool { return player.host == origin }

// Name func
func (player Player) Name() string { return player.name }

// IsNameEmpty func
func (player Player) IsNameEmpty() bool { return player.name == "" }

// IsHandEmpty func
func (player Player) IsHandEmpty() bool { return len(player.hand) == 0 }

// Fold func
func (player *Player) Fold() { player.fold = true }

// Play function
func (player *Player) Play(card card.ID) error {
	index := player.hand.Find(card)
	if index == -1 {
		return errors.New("Card is not in players hand")
	}
	player.hand.Remove(index)
	return nil
}

// Exchange func
func (player *Player) Exchange(card card.ID, side *deck.Cards) error {
	index := player.hand.Find(card)
	if index == -1 {
		return errors.New("Card is not in players hand")
	}
	player.hand.Add((*side)[0])
	side.Remove(0)
	side.Add(card)
	player.hand.Remove(index)
	return nil
}

// Collect func
func (player *Player) Collect(cards *deck.Cards) {
	player.pile.Add(*cards...)
}

// Points func
func (player Player) Points(scorer func(card.ID) uint8) uint8 {
	return player.pile.Sum(scorer)
}

// IsExpectedPlayer func
func (player *Player) IsExpectedPlayer(other *Player, origin string) bool {
	return player == other && player.IsSameHost(origin)
}

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %v)",
		player.name, player.hand, player.pile, player.fold)
}

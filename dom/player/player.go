package player

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

// Player struct
type Player struct {
	name, host string
	hand, pile deck.Cards
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

// Hand func
func (player *Player) Hand() *deck.Cards {
	return &player.hand
}

// Pile func
func (player *Player) Pile() *deck.Cards {
	return &player.pile
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

// IsSameHost func
func (player Player) IsSameHost(origin string) bool { return player.host == origin }

// Name func
func (player Player) Name() string { return player.name }

// Predicate type
type Predicate func(p *Player) bool

// IsNameEmpty func
var IsNameEmpty Predicate = func(p *Player) bool { return p.name == "" }

// IsHandEmpty func
var IsHandEmpty Predicate = func(p *Player) bool { return len(p.hand) == 0 }

// Folded func
var Folded Predicate = func(p *Player) bool { return p.fold }

// Fold func
func (player *Player) Fold() { player.fold = true }

// Collect func
func (player *Player) Collect(cards *deck.Cards) {
	player.pile.Add(*cards...)
}

// Points func
func (player Player) Points(scorer func(card.ID) uint8) uint8 {
	return player.pile.Sum(scorer)
}

func (player Player) String() string {
	return fmt.Sprintf("(Name: %s, Cards: %+v, Pile: %+v, Has folded? %t)",
		player.name, player.hand, player.pile, player.fold)
}

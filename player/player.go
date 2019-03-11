package player

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/point"
)

// Player struct
type Player struct {
	name         string
	host         string
	hand         deck.Cards
	pile         deck.Cards
	auctionScore uint8
	fold         bool
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
func (player *Player) Folded() bool {
	return player.fold
}

// IsRemoteHost func
func (player *Player) IsRemoteHost(origin string) bool {
	return player.host == origin
}

// IsName func
func (player *Player) IsName(name string) bool {
	return player.name == name
}

// Fold func
func (player *Player) Fold() {
	player.fold = true
}

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
func (player *Player) Count() uint8 {
	return point.Count(player.pile, briscola.Points)
}

func (player Player) String() string {
	host := display.NewInfo("Host", ":", player.host, ";")
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	pile := display.NewInfo("Pile", ":", player.pile.String(), ";")
	fold := display.NewInfo("Folded", ":", strconv.FormatBool(player.Folded()), ";")
	return display.All(display.Wrap("Player", player.Name(), host, hand, pile, fold)...)
}

// Name func
func (player *Player) Name() display.Info {
	return display.NewInfo("Name", ":", player.name, ";")
}

// Info function
func (player Player) Info() []display.Info {
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	return display.Wrap("Player", player.Name(), hand)
}

package player

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/display"
)

// Player struct
type Player struct {
	info         info
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
	player.info = info{name, origin}
}

// Folded func
func (player *Player) Folded() bool {
	return player.fold
}

// NotFolded func
func (player *Player) NotFolded() bool {
	return !player.fold
}

// IsSame func
func (player *Player) IsSame(other *Player) bool {
	return player == other
}

// IsSameHost func
func (player *Player) IsSameHost(origin string) bool {
	return player.info.IsSameHost(origin)
}

// IsName func
func (player *Player) IsName(name string) bool {
	return player.info.IsName(name)
}

// IsNameEmpty func
func (player *Player) IsNameEmpty() bool { return player.IsName("") }

// IsHandEmpty func
func (player *Player) IsHandEmpty() bool { return len(*player.Hand()) == 0 }

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
func (player *Player) Count(scorer func(card.ID) uint8) uint8 {
	return player.pile.Sum(scorer)
}

// Name func
func (player *Player) Name() display.Info {
	return display.NewInfo("Name", ":", player.info.name, ";")
}

// Info function
func (player Player) Info() []display.Info {
	info := player.info.Info()
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	info = append(info, hand)
	return display.Wrap("Player", info...)
}

func (player Player) String() string {
	info := display.NewInfo("Hand", ":", player.info.String(), ";")
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	pile := display.NewInfo("Pile", ":", player.pile.String(), ";")
	fold := display.NewInfo("Folded", ":", strconv.FormatBool(player.Folded()), ";")
	return display.All(display.Wrap("Player", info, hand, pile, fold)...)
}

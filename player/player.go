package player

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/display"
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
	player.Hand().Add(cards.Supply())
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

// Name func
func (player *Player) Name() string {
	return player.name
}

// IsRemoteHost func
func (player *Player) IsRemoteHost(origin string) bool {
	return player.host == origin
}

// Has func
func (player *Player) Has(id card.ID) bool {
	return player.Hand().Has(id)
}

// Pile func
func (player *Player) Pile() *deck.Cards {
	return &player.pile
}

func (player *Player) collect(cards deck.Cards) {
	if len(cards) > 0 {
		player.Pile().Add(cards...)
	}
}

// Fold func
func (player *Player) Fold() {
	player.fold = true
}

// Folded func
func (player *Player) Folded() bool {
	return player.fold
}

// Play function
func (player *Player) Play(number, seed string) (card.ID, error) {
	inputCard, err := card.Create(number, seed)
	index, err := player.Hand().Find(func(c card.ID) bool { return c == inputCard })
	if err == nil {
		player.Hand().Remove(index)
		return inputCard, nil
	}
	return 0, err
}

func (player Player) String() string {
	name := display.NewInfo("Name", ":", player.Name(), ";")
	host := display.NewInfo("Host", ":", player.host, ";")
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	pile := display.NewInfo("Pile", ":", player.pile.String(), ";")
	fold := display.NewInfo("Folded", ":", strconv.FormatBool(player.Folded()), ";")
	return display.All(display.Wrap("Player", name, host, hand, pile, fold)...)
}

// Info function
func (player Player) Info() []display.Info {
	name := display.NewInfo("Name", ":", player.Name(), ";")
	hand := display.NewInfo("Hand", ":", player.hand.String(), ";")
	return display.Wrap("Player", name, hand)
}

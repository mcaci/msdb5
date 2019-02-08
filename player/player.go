package player

import "github.com/nikiforosFreespirit/msdb5/card"

// Player struct
type Player struct {
	name         string
	host         string
	hand         deck.Cards
	pile         deck.Cards
	auctionScore uint8
}

// New func
func New() *Player {
	player := new(Player)
	player.hand = deck.Cards{}
	return player
}

// Draw func
func (player *Player) Draw(cards deck.Cards) card.ID {
	c := cards.Supply()
	player.Hand().Add(c)
	return c
}

// Hand func
func (player *Player) Hand() *deck.Cards {
	return &player.hand
}

// Name func
func (player *Player) Name() string {
	return player.name
}

// SetName func
func (player *Player) SetName(name string) {
	player.name = name
}

// MyHostIs func
func (player *Player) MyHostIs(host string) {
	player.host = host
}

// Host func
func (player *Player) Host() string {
	return player.host
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

// SetAuctionScore func
func (player *Player) SetAuctionScore(auctionScore uint8) {
	player.auctionScore = auctionScore
}

// AuctionScore func
func (player *Player) AuctionScore() uint8 {
	return player.auctionScore
}

func (player *Player) score(count func(cards deck.Cards) uint8) uint8 {
	return count(*player.Pile())
}

// Play function
func (player *Player) Play(number, seed string) (card.ID, bool) {
	card, _ := card.ByName(number, seed)
	found := false
	for index, c := range *(player.Hand()) {
		found = c == card
		if found {
			player.Hand().Remove(index)
			break
		}
	}
	return card, found
}

func (player Player) String() string {
	str := "Player["
	str += print("Name", player.name)
	str += print("Host", player.host)
	str += print("Hand", player.hand.String())
	str += print("Pile", player.pile.String())
	str += "]"
	return str
}

func print(info, field string) string {
	return info + ":" + field + ";"
}

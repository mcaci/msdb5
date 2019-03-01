package player

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

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
func (player *Player) Draw(cards deck.Cards) {
	player.Hand().Add(cards.Supply())
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

func (player Player) head() string {
	return "Player"
}

func (player Player) tail() string {
	return ""
}

func (player Player) handInfo() string {
	return player.hand.String()
}

func (player Player) pileInfo() string {
	return player.pile.String()
}

func (player Player) auctionScoreInfo() string {
	return strconv.Itoa(int(player.auctionScore))
}

func (player Player) String() (str string) {
	str += print("Type", player.head())
	str += print("Name", player.Name())
	str += print("Host", player.Host())
	str += print("Hand", player.handInfo())
	str += print("Pile", player.pileInfo())
	str += print("AuctionScore", player.auctionScoreInfo())
	str += print("End", player.tail())
	return
}

// Print function
func (player Player) Print() (str string) {
	str += print("Type", player.head())
	str += print("Name", player.Name())
	str += print("Hand", player.handInfo())
	if player.auctionScore > 0 {
		str += print("AuctionScore", player.auctionScoreInfo())
	}
	str += print("End", player.tail())
	return
}

func print(info, field string) string {
	return info + ":" + field + ";"
}

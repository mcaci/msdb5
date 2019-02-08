package board

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/api"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	players      player.Players
	playedCards  deck.Cards
	selectedCard card.ID
	auctionScore uint8
}

// New func
func New() *Board {
	b := new(Board)
	makePlayers(b)
	playersDrawAllCards(&b.players)
	return b
}

// NewAction func
func NewAction() api.Action {
	return New()
}

func makePlayers(b *Board) {
	for i := 0; i < 5; i++ {
		b.players.Add(*player.New())
	}
}

func playersDrawAllCards(players *player.Players) {
	deck := card.Deck()
	for i := 0; i < card.DeckSize; i++ {
		(*players)[i%5].Draw(deck)
	}
}

// Players func
func (b *Board) Players() player.Players {
	return b.players
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}

// NominatedCard func
func (b *Board) NominatedCard() *card.ID {
	return &b.selectedCard
}

// PlayedCards func
func (b *Board) PlayedCards() *deck.Cards {
	return &b.playedCards
}

// String func
func (b Board) String() string {
	var str string
	str += "Board("
	str += "Players[" + b.players.String() + "]"
	str += "PlayedCards[" + b.playedCards.String() + "]"
	str += "SelectedCard[" + b.selectedCard.String() + "]"
	str += "AuctionScore[" + strconv.Itoa(int(b.auctionScore)) + "]"
	str += ")"
	return str
}

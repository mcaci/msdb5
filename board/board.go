package board

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Board struct
type Board struct {
	playedCards  deck.Cards
	auctionScore uint8
}

// New func
func New() *Board {
	return new(Board)
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}

// PlayedCards func
func (b *Board) PlayedCards() *deck.Cards {
	return &b.playedCards
}

// Print func
func (b Board) Print() string {
	var str string
	str += "Board("
	str += "PlayedCards[" + b.playedCards.String() + "]"
	str += "AuctionScore[" + strconv.Itoa(int(b.auctionScore)) + "]"
	str += ")"
	return str
}

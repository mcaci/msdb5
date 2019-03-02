package board

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/display"
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

// PlayedCardIs func
func (b *Board) PlayedCardIs(card card.ID) bool {
	b.PlayedCards().Add(card)
	return len(*b.PlayedCards()) >= 5
}

// Info func
func (b Board) Info() []display.Info {
	pCar := display.NewInfo("PlayedCards", ":", b.playedCards.String(), ";")
	aSco := display.NewInfo("AuctionScore", ":", strconv.Itoa(int(b.auctionScore)), ";")
	return display.Wrap("Board", pCar, aSco)
}

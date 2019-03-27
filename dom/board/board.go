package board

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

// Board struct
type Board struct {
	side         deck.Cards
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

// SideDeck func
func (b *Board) SideDeck() *deck.Cards {
	return &b.side
}

func (b Board) String() string {
	return fmt.Sprintf("(Played cards: %+v, Side deck: %+v, Auction score: %d)",
		b.playedCards, b.side, b.auctionScore)
}

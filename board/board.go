package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	players      []*player.Player
	playedCards  card.Cards
	auctionScore uint8
}

// New func
func New() *Board {
	var b Board

	b.players = make([]*player.Player, 5)
	for i := range b.players {
		b.players[i] = player.New()
	}

	deck := card.Deck()
	for i := 0; i < card.DeckSize; i++ {
		b.players[i%5].Draw(&deck)
	}

	b.playedCards = card.Cards{}

	return &b
}

// Players func
func (b *Board) Players() []*player.Player {
	return b.players
}

// PlayedCards func
func (b *Board) PlayedCards() *card.Cards {
	return &b.playedCards
}

// SetAuctionScore func
func (b *Board) SetAuctionScore(score uint8) {
	b.auctionScore = score
}

// AuctionScore func
func (b *Board) AuctionScore() uint8 {
	return b.auctionScore
}

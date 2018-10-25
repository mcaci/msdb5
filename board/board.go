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
	b := new(Board)
	makePlayers(b)
	playersDrawAllCards(&b.players)
	return b
}

func makePlayers(b *Board) {
	for i := 0; i < 5; i++ {
		b.players = append(b.players, player.New())
	}
}

func playersDrawAllCards(players *[]*player.Player) {
	deck := card.Deck()
	for i := 0; i < card.DeckSize; i++ {
		(*players)[i%5].Draw(&deck)
	}
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

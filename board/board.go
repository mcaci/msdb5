package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	deck        card.Cards
	players     []*player.Player
	playedCards card.Cards
}

// New func
func New() *Board {
	var b Board

	b.deck = card.Deck()

	b.players = make([]*player.Player, 5)
	for i := range b.players {
		b.players[i] = player.New()
	}
	for i := 0; i < card.DeckSize; i++ {
		b.players[i%5].Draw(&b.deck)
	}

	b.playedCards = card.Cards{}

	return &b
}

// Deck func
func (b *Board) Deck() card.Cards {
	return b.deck
}

// Players func
func (b *Board) Players() []*player.Player {
	return b.players
}

// PlayedCards func
func (b *Board) PlayedCards() *card.Cards {
	return &b.playedCards
}

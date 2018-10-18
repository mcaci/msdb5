package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	deck        deck.Cards
	players     []player.Player
	playedCards card.Cards
}

// New func
func New() *Board {
	var b Board

	b.deck = deck.New()

	b.players = make([]player.Player, 5)
	for i := range b.players {
		b.players[i] = player.New()
	}
	for i := 0; !b.deck.IsEmpty(); i++ {
		b.players[i%5].Draw(&b.deck)
	}

	b.playedCards = card.Cards{}

	return &b
}

// Deck func
func (b *Board) Deck() deck.Cards {
	return b.deck
}

// Players func
func (b *Board) Players() []player.Player {
	return b.players
}

// PlayedCards func
func (b *Board) PlayedCards() *card.Cards {
	return &b.playedCards
}

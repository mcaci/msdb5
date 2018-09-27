package board

import (
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	deck    deck.Deck
	players []player.Player
}

// New func
func New() *Board {
	var b Board

	b.deck = deck.New()

	b.players = make([]player.Player, 5)
	for i := range b.players {
		b.players[i] = player.New()
	}

	return &b
}

// Deck func
func (b *Board) Deck() deck.Deck {
	return b.deck
}

// Players func
func (b *Board) Players() []player.Player {
	return b.players
}

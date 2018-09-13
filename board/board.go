package board

import (
	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Board struct
type Board struct {
	deck deck.Deck
}

// New func
func New() *Board {
	var b Board
	b.deck = deck.New()
	return &b
}

// Deck func
func (b *Board) Deck() deck.Deck {
	return b.deck
}

func (b *Board) Player() player.Player {
	return nil
}
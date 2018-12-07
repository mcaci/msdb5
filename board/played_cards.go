package board

import (
	"github.com/nikiforosFreespirit/msdb5/card/set"
)

// PlayedCards func
func (b *Board) PlayedCards() *set.Cards {
	return &b.playedCards
}

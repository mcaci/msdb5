package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

// PlayedCards func
func (b *Board) PlayedCards() *set.Cards {
	return &b.playedCards
}

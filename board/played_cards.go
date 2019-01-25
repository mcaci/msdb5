package board

import "github.com/nikiforosFreespirit/msdb5/card"

// PlayedCards func
func (b *Board) PlayedCards() *card.Cards {
	return &b.playedCards
}

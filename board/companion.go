package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Nominate func
func (b *Board) Nominate(number, seed, origin string) (card.ID, error) {
	card, err := card.ByName(number, seed)
	if err == nil {
		b.selectedCard = card
	}
	return card, err
}

// NominatedCard func
func (b *Board) NominatedCard() *card.ID {
	return &b.selectedCard
}

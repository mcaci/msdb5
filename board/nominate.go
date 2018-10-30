package board

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Nominate func
func (b *Board) Nominate(number, seed string) (card.Card, error) {
	card, err := card.ByName(number, seed)
	if err == nil {
		b.selectedCard = card
	}
	return card, err
}

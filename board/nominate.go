package board

import (
	"github.com/nikiforosFreespirit/msdb5/board/prompt"
	"github.com/nikiforosFreespirit/msdb5/card"
)

// AskNominatedCard func
func (b *Board) AskNominatedCard(promptFunc func(chan<- card.ID), playerToAsk chan card.ID) card.ID {
	b.selectedCard = prompt.Card(promptFunc, playerToAsk)
	return b.selectedCard
}

// Nominate func
func (b *Board) Nominate(number, seed string) (card.ID, error) {
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

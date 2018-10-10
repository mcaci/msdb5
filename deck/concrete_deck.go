package deck

import "github.com/nikiforosFreespirit/msdb5/card"

type concreteDeck struct {
	cards []int
	index int
}

// IsEmpty func
func (deck *concreteDeck) IsEmpty() bool {
	return deck.index >= 40 // should be 40 as it's the deck's size
}

// RemoveTop func
func (deck *concreteDeck) RemoveTop() card.Card {
	index := deck.index
	deck.index++
	card, err := card.ByID(uint8(deck.cards[index] + 1))
	if err != nil {
		panic("should not be here")
	}
	return card
}

package deck

import "github.com/nikiforosFreespirit/msdb5/card"

type concreteDeck struct {
	cards []int
	index int
}

// IsEmpty func
func (deck *concreteDeck) IsEmpty() bool {
	return deck.index >= Size // should be 40 as it's the deck's size
}

// Supply func
func (deck *concreteDeck) Supply() card.Card {
	index := deck.index
	deck.index++
	card, err := card.ByID(uint8(deck.cards[index] + 1))
	if err != nil {
		panic("should not be here")
	}
	return card
}

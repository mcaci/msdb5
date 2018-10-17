package deck

import "github.com/nikiforosFreespirit/msdb5/card"

type concreteDeck struct {
	cards []int
	index int
}

// Get func
func (deck *concreteDeck) Get() []int {
	return deck.cards
}

// IsEmpty func
func (deck *concreteDeck) IsEmpty() bool {
	return deck.index >= Size // should be 40 as it's the deck's size
}

// Supply func
func (deck *concreteDeck) Supply() card.Card {
	card, err := card.ByID(deck.topCardID())
	if err != nil {
		panic("Should not be here: " + err.Error())
	} else {
		deck.updateTopCardIndex()
	}
	return card
}

func (deck *concreteDeck) updateTopCardIndex() {
	deck.index++
}

func (deck *concreteDeck) topCardID() uint8 {
	return uint8(deck.cards[deck.index] + 1)
}

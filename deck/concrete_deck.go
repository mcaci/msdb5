package deck

import "github.com/nikiforosFreespirit/msdb5/card"

type concreteDeck struct {
	cards []uint8
}

func (deck *concreteDeck) GetIDs() []uint8 {
	return deck.cards
}

// IsEmpty func
func (deck *concreteDeck) IsEmpty() bool {
	return len(deck.cards) <= 0 // should be 40 as it's the deck's size
}

// Supply func
func (deck *concreteDeck) Supply() card.Card {
	card, err := card.ByID(deck.cards[0])
	if err != nil {
		panic("Should not be here: " + err.Error())
	} else {
		deck.updateDeck()
	}
	return card
}

func (deck *concreteDeck) updateDeck() {
	deck.cards = deck.cards[1:]
}

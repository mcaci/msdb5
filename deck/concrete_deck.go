package deck

import "github.com/nikiforosFreespirit/msdb5/card"

type concreteDeck struct {
	cards []int
	index int
}

// Get func
func (deck *concreteDeck) GetIDs() []uint8 {
	ids := []uint8{}
	for index := range deck.cards {
		ids = append(ids, deck.idAt(index))
	}
	return ids
}

// IsEmpty func
func (deck *concreteDeck) IsEmpty() bool {
	return deck.index >= Size // should be 40 as it's the deck's size
}

// Supply func
func (deck *concreteDeck) Supply() card.Card {
	card, err := card.ByID(deck.idAt(deck.index))
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

func (deck *concreteDeck) idAt(index int) uint8 {
	return uint8(deck.cards[index] + 1)
}

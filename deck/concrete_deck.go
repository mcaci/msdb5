package deck

import "github.com/nikiforosFreespirit/msdb5/card"

// Cards type
type Cards []uint8

// GetIDs func
func (deck *Cards) GetIDs() []uint8 {
	return []uint8(*deck)
}

// IsEmpty func
func (deck *Cards) IsEmpty() bool {
	return len(*deck) <= 0 // should be 40 as it's the deck's size
}

// Supply func
func (deck *Cards) Supply() card.Card {
	card, err := card.ByID((*deck)[0])
	if err != nil {
		panic("Should not be here: " + err.Error())
	} else {
		deck.updateDeck()
	}
	return card
}

func (deck *Cards) updateDeck() {
	(*deck) = (*deck)[1:]
}

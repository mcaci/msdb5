package card

import (
	"math/rand"
	"time"
)

// Cards type
type Cards []ID

// Add func
func (cards *Cards) Add(id ID) {
	*cards = append(*cards, id)
}

// Has func
func (cards Cards) Has(id ID) bool {
	var found bool
	for _, cardID := range cards {
		if found = (cardID == id); found {
			break
		}
	}
	return found
}

// IsEmpty func
func (cards *Cards) IsEmpty() bool {
	return len(*cards) <= 0 // should be 40 as it's the cards's size
}

// Supply func
func (cards *Cards) Supply() ID {
	card := (*cards)[0]
	cards.updateDeck()
	return card
}

func (cards *Cards) updateDeck() {
	(*cards) = (*cards)[1:]
}

// DeckSize of a cards of cards
const DeckSize = 40

// Deck func
func Deck() Cards {
	cards := new(Cards)

	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(DeckSize)
	for index := range ints {
		*cards = append(*cards, ID(ints[index]+1))
	}
	return *cards
}

// FillWithIDs func
func FillWithIDs(ids ...ID) Cards {
	var cards Cards
	for _, id := range ids {
		cards.Add(id)
	}
	return cards
}

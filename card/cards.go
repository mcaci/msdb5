package card

import (
	"math/rand"
	"time"
)

// Cards type
type Cards []ID

// Add func
func (cards *Cards) Add(ids ...ID) {
	*cards = append(*cards, ids...)
}

// FillWithIDs func
func FillWithIDs(ids ...ID) Cards {
	var cards Cards
	cards = append(cards, ids...)
	return cards
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

// Supply func
func (cards *Cards) Supply() ID {
	card := (*cards)[0]
	(*cards) = (*cards)[1:]
	return card
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

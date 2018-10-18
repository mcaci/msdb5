package card

import (
	"math/rand"
	"time"
)

// Cards type
type Cards []uint8

// Add func
func (cards *Cards) Add(id uint8) {
	*cards = append(*cards, id)
}

// Has func
func (cards Cards) Has(id uint8) bool {
	var cardFound bool
	for _, c := range cards {
		cardFound = (c == id)
		if cardFound {
			break
		}
	}
	return cardFound
}

// IsEmpty func
func (cards *Cards) IsEmpty() bool {
	return len(*cards) <= 0 // should be 40 as it's the cards's size
}

// Supply func
func (cards *Cards) Supply() uint8 {
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
		*cards = append(*cards, uint8(ints[index]+1))
	}
	return *cards
}

// FillWithIDs func
func FillWithIDs(ids ...uint8) Cards {
	var cards Cards
	for _, id := range ids {
		cards.Add(id)
	}
	return cards
}

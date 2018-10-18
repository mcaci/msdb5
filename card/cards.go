package card

import (
	"math/rand"
	"time"
)

// Cards type
type Cards []uint8

// Add func
func (cards *Cards) Add(card Card) {
	*cards = append(*cards, card.ID())
}

// Has func
func (cards Cards) Has(card Card) bool {
	var cardFound bool
	for _, c := range cards {
		cardFound = (c == card.ID())
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
func (cards *Cards) Supply() Card {
	card, err := ByID((*cards)[0])
	if err != nil {
		panic("Should not be here: " + err.Error())
	} else {
		cards.updateDeck()
	}
	return card
}

func (cards *Cards) updateDeck() {
	(*cards) = (*cards)[1:]
}

// Size of a cards of cards
const Size = 40

// Deck func
func Deck() Cards {
	cards := new(Cards)

	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(Size)
	for index := range ints {
		*cards = append(*cards, uint8(ints[index]+1))
	}
	return *cards
}

// FillWithIDs func
func FillWithIDs(ids ...uint8) Cards {
	var cards Cards
	for _, id := range ids {
		card, _ := ByID(id)
		cards.Add(card)
	}
	return cards
}

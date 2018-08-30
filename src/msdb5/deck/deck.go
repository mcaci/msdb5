package deck

import (
	"math/rand"
	"msdb5/card"
	"time"
)

// Deck interface
type Deck interface {
	RemoveTop() *(card.Card)
	IsEmpty() bool
}

// ConcreteDeck type
type ConcreteDeck struct {
	cards []int
	index int
}

// Create func
func (deck *ConcreteDeck) Create() {
	const deckSize = 40
	rand.Seed(time.Now().UnixNano())
	deck.cards = rand.Perm(deckSize)
}

// IsEmpty func
func (deck *ConcreteDeck) IsEmpty() bool {
	return deck.index >= 40 // should be 40 as it's the deck's size
}

// RemoveTop func
func (deck *ConcreteDeck) RemoveTop() *(card.Card) {
	index := deck.index
	deck.index++
	card, err := card.ByID(uint8(deck.cards[index] + 1))
	if err != nil {
		panic("should not be here")
	}
	return card
}

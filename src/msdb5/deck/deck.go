package deck

import (
	"math/rand"
	"msdb5/card"
)

// Deck interface
type Deck interface {
	RemoveTop() *(card.Card)
}

// ConcreteDeck type
type ConcreteDeck struct {
	cards []int
	index int
}

// Create func
func (deck *ConcreteDeck) Create() {
	const deckSize = 40
	deck.cards = rand.Perm(deckSize)
}

// RemoveTop func
func (deck *ConcreteDeck) RemoveTop() *(card.Card) {
	index := deck.index
	deck.index++
	return card.ByID(deck.cards[index])
}

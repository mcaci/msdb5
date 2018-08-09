package msdb5

import (
	"math/rand"
)

// Deck interface
type Deck interface {
	RemoveTop() *Card
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
func (deck *ConcreteDeck) RemoveTop() *Card {
	index := deck.index
	deck.index++
	return CardByID(deck.cards[index])
}

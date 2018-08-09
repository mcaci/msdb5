package msdb5

import (
	"math/rand"
	"msdb5/card"
)

// Card : Importing struct from card package
type Card card.Card

// CardPtr : Importing struct pointer from card package
type CardPtr *(card.Card)

// Deck interface
type Deck interface {
	RemoveTop() CardPtr
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
func (deck *ConcreteDeck) RemoveTop() CardPtr {
	index := deck.index
	deck.index++
	return card.ByID(deck.cards[index])
}

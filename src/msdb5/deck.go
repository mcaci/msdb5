package msdb5

import (
	"math/rand"
)

type Deck interface {
	RemoveTop() *Card
}

type ConcreteDeck struct {
	cards []int
	index int
}

func (deck *ConcreteDeck) Create() {
	const deckSize = 40
	deck.cards = rand.Perm(deckSize)
}

func (deck *ConcreteDeck) RemoveTop() *Card {
	index := deck.index
	deck.index++
	return CardById(deck.cards[index])
}

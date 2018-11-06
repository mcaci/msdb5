package card

import (
	"math/rand"
	"time"
)

// DeckSize of a cards of cards
const DeckSize = 40

// Deck func
func Deck() Cards {
	var ids Cards
	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(DeckSize)
	for index := range ints {
		card, _ := Card(fromZeroBased(ints[index]))
		ids.Add(card)
	}
	return ids
}
func fromZeroBased(index int) uint8 {
	return uint8(index + 1)
}

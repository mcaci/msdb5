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
		ids.Add(fromZeroBased(uint8(ints[index])))
	}
	return ids
}

func fromZeroBased(index uint8) ID {
	return ID(index + 1)
}

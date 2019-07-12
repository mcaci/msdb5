package deck

import (
	"math/rand"
	"time"

	"github.com/mcaci/msdb5/dom/card"
)

// DeckSize of a cards of cards
const DeckSize = 40

// New func
func New() (cards Cards) {
	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(DeckSize)
	for _, cardID := range ints {
		cards.Add(card.ID(fromZeroBased(cardID)))
	}
	return
}

func fromZeroBased(index int) uint8 {
	return uint8(index) + 1
}

package deck

import (
	"math/rand"
	"time"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// DeckSize of a cards of cards
const DeckSize = 40

// Deck func
func Deck() Cards {
	var ids Cards
	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(DeckSize)
	for _, cardID := range ints {
		card, _ := card.Create(mapToValidID(cardID))
		ids.Add(card)
	}
	return ids
}

func mapToValidID(index int) uint8 {
	return uint8(index + 1)
}

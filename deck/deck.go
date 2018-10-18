package deck

import (
	"math/rand"
	"time"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// Size of a deck of cards
const Size = 40

// Deck interface
type Deck interface {
	card.Supplier
	IsEmpty() bool
	GetIDs() []uint8
}

// New func
func New() Deck {
	deck := new(concreteDeck)

	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(Size)
	for index := range ints {
		deck.cards = append(deck.cards, uint8(ints[index]+1))
	}
	return deck
}

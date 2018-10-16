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
	RemoveTop() card.Card
	IsEmpty() bool
}

// New func
func New() Deck {
	deck := new(concreteDeck)

	rand.Seed(time.Now().UnixNano())
	deck.cards = rand.Perm(Size)
	return deck
}

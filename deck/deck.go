package deck

import (
	"math/rand"
	"time"

	"github.com/nikiforosFreespirit/msdb5/card"
)

// Cards type
type Cards []uint8

// GetIDs func
func (cards *Cards) GetIDs() []uint8 {
	return []uint8(*cards)
}

// IsEmpty func
func (cards *Cards) IsEmpty() bool {
	return len(*cards) <= 0 // should be 40 as it's the cards's size
}

// Supply func
func (cards *Cards) Supply() card.Card {
	card, err := card.ByID((*cards)[0])
	if err != nil {
		panic("Should not be here: " + err.Error())
	} else {
		cards.updateDeck()
	}
	return card
}

func (cards *Cards) updateDeck() {
	(*cards) = (*cards)[1:]
}

// Size of a cards of cards
const Size = 40

// Deck func
func Deck() Cards {
	cards := new(Cards)

	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(Size)
	for index := range ints {
		*cards = append(*cards, uint8(ints[index]+1))
	}
	return *cards
}

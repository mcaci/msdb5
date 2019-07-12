package deck

import (
	"github.com/mcaci/msdb5/dom/card"
)

// Cards type
type Cards []card.ID

// Add func
func (cards *Cards) Add(ids ...card.ID) {
	*cards = append(*cards, ids...)
}

// Clear func
func (cards *Cards) Clear() {
	*cards = Cards{}
}

// Sum func
func (cards *Cards) Sum(point func(card.ID) uint8) (sum uint8) {
	for _, c := range *cards {
		sum += point(c)
	}
	return
}

// Find func
func (cards *Cards) Find(id card.ID) int {
	for index, c := range *cards {
		if c == id {
			return index
		}
	}
	return -1
}

// Supply func
func (cards *Cards) Supply() card.ID {
	card := (*cards)[0]
	*cards = (*cards)[1:]
	return card
}

// BriscolaSerie func
func BriscolaSerie(briscola card.Seed) Cards {
	set := Cards{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	if briscola != card.Coin {
		for i := range set {
			set[i] += card.ID(10 * briscola)
		}
	}
	return set
}

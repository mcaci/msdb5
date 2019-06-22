package deck

import (
	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

// Cards type
type Cards []card.ID

// Add func
func (cards *Cards) Add(ids ...card.ID) {
	*cards = append(*cards, ids...)
}

// Remove func
func (cards *Cards) Remove(index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
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
	cards.Remove(0)
	return card
}

// Highest func
func Highest(briscola card.Seed) Cards {
	set := Cards{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	delta := card.ID(0)
	switch briscola {
	case card.Cup:
		delta = 10
	case card.Sword:
		delta = 20
	case card.Cudgel:
		delta = 30
	}
	if delta > 0 {
		for i := range set {
			set[i] += delta
		}
	}
	return set
}

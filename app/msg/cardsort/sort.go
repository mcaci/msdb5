package cardsort

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/briscola"
)

type SortedCard struct {
	cards    set.Cards
	briscola *card.Seed
}

func NewSorted(cards set.Cards, b *card.Seed) *SortedCard {
	return &SortedCard{cards, b}
}

func (ids SortedCard) Len() int { return len(ids.cards) }

func (ids SortedCard) Less(i, j int) bool {
	first, other := ids.cards[i], ids.cards[j]
	isSameSeed := first.Seed() == other.Seed()
	if ids.briscola == nil {
		isOtherGreaterOnPoints := briscola.Points(first) < briscola.Points(other)
		isSamePoints := briscola.Points(first) == briscola.Points(other)
		isOtherGreaterOnNumber := first.Number() < other.Number()
		return !(isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints))
	}
	isOtherGreaterOnPoints := briscola.Points(first) < briscola.Points(other)
	isSamePoints := briscola.Points(first) == briscola.Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return !(isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints))
}

func (ids SortedCard) Swap(i, j int) { ids.cards[i], ids.cards[j] = ids.cards[j], ids.cards[i] }

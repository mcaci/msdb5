package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type SortedCard struct {
	cards    set.Cards
	briscola *card.Seed
}

func NewSorted(cards set.Cards, briscola *card.Seed) *SortedCard {
	return &SortedCard{cards, briscola}
}

func (ids SortedCard) Len() int { return len(ids.cards) }

func (ids SortedCard) Less(i, j int) bool {
	first, other := ids.cards[i], ids.cards[j]
	isSameSeed := first.Seed() == other.Seed()
	if ids.briscola == nil {
		isOtherGreaterOnPoints := Points(first) < Points(other)
		isSamePoints := Points(first) == Points(other)
		isOtherGreaterOnNumber := first.Number() < other.Number()
		return !(isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints))
	}
	isOtherGreaterOnPoints := Points(first) < Points(other)
	isSamePoints := Points(first) == Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return !(isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints))
}

func (ids SortedCard) Swap(i, j int) { ids.cards[i], ids.cards[j] = ids.cards[j], ids.cards[i] }

// Points func
func Points(id card.Item) uint8 {
	var points = map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}
	return points[id.Number()]
}

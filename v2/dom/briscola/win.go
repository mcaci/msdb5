package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func Winner(cardsOnTheTable set.Cards, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	id := 0
	for i, other := range cardsOnTheTable {
		if !otherWins(base, other, briscola) {
			continue
		}
		base = other
		id = i
	}
	return uint8(id)
}

func otherWins(first, other card.Item, briscola card.Seed) bool {
	switch first.Seed() == other.Seed() {
	case true:
		if Points(first) != Points(other) {
			return Points(first) < Points(other)
		}
		return first.Number() < other.Number()
	default:
		// case false:
		return other.Seed() == briscola
	}
}

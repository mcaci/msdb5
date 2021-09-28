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

func points(scorer interface{ Number() uint8 }) uint8 {
	var points = map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}
	return points[scorer.Number()]
}

func otherWins(first, other card.Item, briscola card.Seed) bool {
	switch first.Seed() == other.Seed() {
	case true:
		if points(first) != points(other) {
			return points(first) < points(other)
		}
		return first.Number() < other.Number()
	default:
		// case false:
		return other.Seed() == briscola
	}
}

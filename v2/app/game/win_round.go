package game

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func indexOfWinningCard(cardsOnTheTable set.Cards, b card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if winningCard(base, other, b) == other {
			base = other
			max = i
		}
	}
	return uint8(max)
}

func winningCard(base, other card.Item, b card.Seed) card.Item {
	if &base == nil || doesOtherCardWin(base, other, b) {
		base = other
	}
	return base
}

func doesOtherCardWin(first, other card.Item, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := first.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || isOtherHigher(first, other)
}

func isOtherHigher(first, other card.Item) bool {
	isSameSeed := first.Seed() == other.Seed()
	isOtherGreaterOnPoints := briscola.Points(first) < briscola.Points(other)
	isSamePoints := briscola.Points(first) == briscola.Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}

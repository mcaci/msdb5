package briscola

import (
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
)

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable deck.Cards, briscola seeder) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if winningCard(base, other, briscola) == other {
			base = other
			max = i
		}
	}
	return uint8(max)
}

func winningCard(base, other card.ID, briscola seeder) card.ID {
	if &base == nil || doesOtherCardWin(base, other, briscola) {
		base = other
	}
	return base
}

func doesOtherCardWin(first, other card.ID, briscola seeder) bool {
	otherIsBriscola := other.Seed() == briscola.Seed()
	isSameSeed := first.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || isOtherHigher(first, other)
}

func isOtherHigher(first, other card.ID) bool {
	isSameSeed := first.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(first) < Points(other)
	isSamePoints := Points(first) == Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}

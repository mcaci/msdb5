package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// IndexOfWinningCard func - ADUMP - TESTED
func IndexOfWinningCard(cardsOnTheTable card.Cards, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if WinningCard(base, other, briscola) == other {
			base = other
			max = i
		}
	}
	return uint8(max)
}

// WinningCard func - ADUMP - Used above in IndexOfWinningCard
func WinningCard(base, other card.ID, briscola card.Seed) card.ID {
	if &base == nil || doesOtherCardWin(base, other, briscola) {
		base = other
	}
	return base
}

func doesOtherCardWin(first, other card.ID, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := first.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || isOtherHigher(first, other)
}

func isOtherHigher(first, other card.ID) bool {
	isSameSeed := first.Seed() == other.Seed()
	isOtherGreaterOnPoints := points(first) < points(other)
	isSamePoints := points(first) == points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}

package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// WinningCard func
func WinningCard(base, other card.ID, briscola card.Seed) card.ID {
	if &base == nil || DoesOtherCardWin(base, other, briscola) {
		base = other
	}
	return base
}

// DoesOtherCardWin function
func DoesOtherCardWin(first, other card.ID, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
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

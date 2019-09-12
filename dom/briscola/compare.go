package briscola

import (
	"github.com/mcaci/ita-cards/card"
)

// DoesOtherCardWin func
func DoesOtherCardWin(first, other card.Item, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := first.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || IsOtherHigher(first, other)
}

// IsOtherHigher func
func IsOtherHigher(first, other card.Item) bool {
	isSameSeed := first.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(first) < Points(other)
	isSamePoints := Points(first) == Points(other)
	isOtherGreaterOnNumber := first.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}

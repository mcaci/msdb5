package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// DoesOtherCardWin function
func DoesOtherCardWin(base, other card.ID, briscola card.Seed) bool {
	baseIsNotBriscola := base.Seed() != briscola
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := base.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isOtherGreaterOnNumberOnly := base.Number() < other.Number() && Points(base) == Points(other)
	// onlyOtherCardIsBriscola := baseIsNotBriscola && otherIsBriscola
	// otherCardIsBiggerAndOfSameSeed := !isSameSeed && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
	return (baseIsNotBriscola && otherIsBriscola) || (isSameSeed && isOtherGreaterOnPoints) || (isSameSeed && isOtherGreaterOnNumberOnly)
}

func doesOtherCardWin(base, other card.ID) bool {
	isSameSeed := base.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isOtherGreaterOnNumberOnly := Points(base) == Points(other) && base.Number() < other.Number()
	return isSameSeed && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
}

func isBriscola(card card.ID, briscola card.Seed) bool {
	return card.Seed() == briscola
}

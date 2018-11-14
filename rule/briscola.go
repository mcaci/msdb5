package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// DoesOtherCardWin function
func DoesOtherCardWin(base, other card.ID, briscola card.Seed) bool {
	baseIsNotBriscola := base.Seed() != briscola
	otherIsBriscola := other.Seed() == briscola
	onlyOtherCardIsBriscola := baseIsNotBriscola && otherIsBriscola
	areSeedDifferent := base.Seed() != other.Seed()
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isOtherGreaterOnNumberOnly := Points(base) == Points(other) && base.Number() < other.Number()
	otherCardIsBiggerAndOfSameSeed := !areSeedDifferent && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
	return onlyOtherCardIsBriscola || otherCardIsBiggerAndOfSameSeed
}

func doesOtherCardWin(base, other card.ID) bool {
	areSeedDifferent := base.Seed() != other.Seed()
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isOtherGreaterOnNumberOnly := Points(base) == Points(other) && base.Number() < other.Number()
	return !areSeedDifferent && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
}

func isBriscola(card card.ID, briscola card.Seed) bool {
	return card.Seed() == briscola
}
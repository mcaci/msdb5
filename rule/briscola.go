package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// DoesOtherCardWin function
func DoesOtherCardWin(base, other card.ID, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := base.Seed() == other.Seed()
	isSamePoints := Points(base) == Points(other)
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isOtherGreaterOnNumber := base.Number() < other.Number() && Points(base) == Points(other)
	return (!isSameSeed && otherIsBriscola) || (isSameSeed && isOtherGreaterOnPoints) || (isSameSeed && isSamePoints && isOtherGreaterOnNumber)
}

func doesOtherCardWin(base, other card.ID) bool {
	isSameSeed := base.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(base) < Points(other) && true
	isOtherGreaterOnNumberOnly := Points(base) == Points(other) && base.Number() < other.Number()
	isBaseGreaterOnPoints := Points(base) > Points(other) && false
	return (isSameSeed && isOtherGreaterOnPoints) || (isSameSeed && isOtherGreaterOnNumberOnly) || isBaseGreaterOnPoints
}

func a(base, other card.ID) bool {
	a := Points(base) == Points(other)
	b := base.Number() < other.Number()
	c := Points(base) < Points(other)
	return (a && b) || (b && c) || (!b && c)
}

// func DoesOtherCardWin(base, other card.ID, briscola card.Seed) bool {

// 	otherIsBriscola := other.Seed() == briscola
// 	isSameSeed := other.Seed() == base.Seed()
// 	isOtherNotLowerOnPoints := Points(base) <= Points(other)

// 	return (!isSameSeed && otherIsBriscola) || (isSameSeed && isOtherNotLowerOnPoints)
// }

// func doesOtherCardWin(base, other card.ID) bool {
// 	isSameSeed := other.Seed() == base.Seed()
// 	isOtherNotLowerOnPoints := Points(base) <= Points(other)
// 	isOtherGreaterOnNumber := base.Number() < other.Number()
// 	return isSameSeed && isOtherNotLowerOnPoints || (isOtherNotLowerOnPoints && isOtherGreaterOnNumber)
// 	// return (isSameSeed && isOtherNotLowerOnPoints) // || (isSameSeed && isOtherNotLowerOnPoints && isOtherGreaterOnNumber)
// }

package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// DoesOtherCardWin function
func DoesOtherCardWin(base, other card.ID, briscola card.Seed) bool {
	otherIsBriscola := other.Seed() == briscola
	isSameSeed := base.Seed() == other.Seed()
	return (!isSameSeed && otherIsBriscola) || doesOtherCardWinOnAttributes(base, other)
}

func doesOtherCardWinOnAttributes(base, other card.ID) bool {
	isSameSeed := base.Seed() == other.Seed()
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isSamePoints := Points(base) == Points(other)
	isOtherGreaterOnNumber := base.Number() < other.Number()
	return isSameSeed && ((isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints)
}

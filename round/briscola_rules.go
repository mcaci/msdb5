package round

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// DoesOtherCardWin function
func DoesOtherCardWin(base, other *card.Card) bool {
	areSeedDifferent := base.Seed() != other.Seed()
	isOtherGreaterOnPoints := base.Points() < other.Points()
	isOtherGreaterOnNumberOnly := base.Points() == other.Points() && base.Number() < other.Number()
	return !areSeedDifferent && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
}

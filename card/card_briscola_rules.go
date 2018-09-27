package card

// DoesOtherCardWin function
func DoesOtherCardWin(base, other *Card) bool {
	areSeedDifferent := base.seed != other.seed
	isOtherGreaterOnPoints := base.Points() < other.Points()
	isOtherGreaterOnNumberOnly := base.Points() == other.Points() && base.number < other.number
	return !areSeedDifferent && (isOtherGreaterOnPoints || isOtherGreaterOnNumberOnly)
}

package auction

// Score is the auction value
type Score uint8

// Update func
func Update(actual, proposed Score) Score {
	const minScore = 61
	const maxScore = 120
	actualScore := proposed
	switch {
	case proposed < actual:
		actualScore = actual
	case proposed < minScore:
		actualScore = minScore
	case proposed > maxScore:
		actualScore = maxScore
	default:
	}
	return actualScore
}

// CheckScores func
func CheckScores(actual, proposed Score) bool {
	return proposed > actual
}

type cmpInfo int8

const (
	LT_MIN_SCORE cmpInfo = iota - 2
	LE_ACTUAL
	GT_ACTUAL
	GE_MAX_SCORE
)

const (
	MIN_SCORE = 61
	MAX_SCORE = 120
)

// CmpAndSet func
func CmpAndSet(actual, proposed Score) Score {
	switch Cmp(actual, proposed) {
	case LT_MIN_SCORE:
		return MIN_SCORE
	case LE_ACTUAL:
		return actual
	case GE_MAX_SCORE:
		return MAX_SCORE
	default:
		return proposed
	}
}

// Cmp func
func Cmp(actual, proposed Score) cmpInfo {
	switch {
	case proposed <= actual:
		return LE_ACTUAL
	case proposed < MIN_SCORE:
		return LT_MIN_SCORE
	case proposed >= MAX_SCORE:
		return GE_MAX_SCORE
	default:
		return GT_ACTUAL
	}
}

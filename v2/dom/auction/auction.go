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

// Max120 func
func Max120(actual, proposed Score) Score {
	const minScore = 61
	const maxScore = 120
	switch {
	case proposed < actual:
		return actual
	case proposed < minScore:
		return minScore
	case proposed > maxScore:
		return maxScore
	default:
		return proposed
	}
}

// CheckScores func
func CheckScores(actual, proposed Score) bool {
	return proposed > actual
}

// ScoreCmp func
func ScoreCmp(actual, proposed Score) int {
	switch {
	case actual < proposed:
		return -1
	case actual > proposed:
		return 1
	default:
		return 0
	}
}

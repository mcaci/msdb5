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

package auction

// Score is the auction value
type Score uint8

// Update func
func (actual *Score) Update(proposed Score) {
	const minScore = 61
	const maxScore = 120
	actualScore := proposed
	switch {
	case proposed < *actual:
		actualScore = *actual
	case proposed < minScore:
		actualScore = minScore
	case proposed > maxScore:
		actualScore = maxScore
	default:
	}
	*actual = actualScore
}

// CheckWith func
func (actual *Score) CheckWith(proposed Score) bool {
	return proposed > *actual
}

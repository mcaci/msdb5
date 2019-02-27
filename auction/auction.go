package auction

// Update func
func Update(baseScore, prevScore, currentScore uint8, set func(uint8)) {
	const minScore = 61
	const maxScore = 120
	actualScore := currentScore
	if currentScore < prevScore {
		actualScore = baseScore
	} else if currentScore < minScore {
		actualScore = minScore
	} else if currentScore > maxScore {
		actualScore = maxScore
	}
	set(actualScore)
}

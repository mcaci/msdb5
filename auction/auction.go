package auction

import (
	"strconv"
)

// Update func
func Update(prevScore, currentScore uint8, set func(uint8)) {
	const minScore = 61
	const maxScore = 120
	actualScore := currentScore
	if currentScore < prevScore {
		actualScore = prevScore
	} else if currentScore < minScore {
		actualScore = minScore
	} else if currentScore > maxScore {
		actualScore = maxScore
	}
	set(actualScore)
}

// CheckAndUpdate func
func CheckAndUpdate(score string, folded func() bool, fold func(), get func() uint8, set func(uint8)) {
	if !folded() {
		prevScore := get()
		currentScore, err := strconv.Atoi(score)
		if err != nil || uint8(currentScore) <= prevScore {
			fold()
		} else {
			Update(prevScore, uint8(currentScore), set)
		}
	}
}

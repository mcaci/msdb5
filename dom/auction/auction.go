package auction

import (
	"strconv"
)

// Score is the auction value
type Score uint8

// Update func
func Update(prevScore, currentScore Score, set func(Score)) {
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
func CheckAndUpdate(score string, folded func() bool, fold func(), get func() Score, set func(Score)) {
	if !folded() {
		prevScore := get()
		currentScore, err := strconv.Atoi(score)
		if err != nil || Score(currentScore) <= prevScore {
			fold()
		} else {
			Update(prevScore, Score(currentScore), set)
		}
	}
}

package score

import "github.com/nikiforosFreespirit/msdb5/card/set"

// Compute func
func Compute(cards set.Cards) (sum uint8) {
	for _, card := range cards {
		sum += card.Points()
	}
	return sum
}

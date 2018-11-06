package score

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Compute func
func Compute(cards card.Cards) (sum uint8) {
	for _, card := range cards {
		sum += card.Points()
	}
	return sum
}

package score

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Compute func
func Compute(cards ...uint8) uint8 {
	var sum uint8
	for _, id := range cards {
		card, _ := card.ByID(id)
		sum += card.Points()
	}
	return sum
}

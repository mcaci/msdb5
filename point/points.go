package point

import (
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Count func
func Count(cards deck.Cards) (sum uint8) {
	for _, card := range cards {
		sum += card.Points()
	}
	return sum
}

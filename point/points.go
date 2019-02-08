package point

import (
	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Count func
func Count(cards deck.Cards) (sum uint8) {
	for _, c := range cards {
		sum += briscola.Points(c)
	}
	return sum
}

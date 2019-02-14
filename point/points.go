package point

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

// Count func
func Count(cards deck.Cards, point func(card.ID) uint8) (sum uint8) {
	for _, c := range cards {
		sum += point(c)
	}
	return sum
}

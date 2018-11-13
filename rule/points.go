package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
)

// Count func
func Count(cards set.Cards) (sum uint8) {
	for _, card := range cards {
		sum += Points(card)
	}
	return sum
}

// Points func
func Points(card card.ID) uint8 {
	switch card.ToNumber() {
	case 1:
		return 11
	case 3:
		return 10
	case 8:
		return 2
	case 9:
		return 3
	case 10:
		return 4
	default:
		return 0
	}
}

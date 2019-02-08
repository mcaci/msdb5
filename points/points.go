package point

import (
	"github.com/nikiforosFreespirit/msdb5/card"
)

// Count func
func Count(cards deck.Cards) (sum uint8) {
	for _, card := range cards {
		sum += points(card)
	}
	return sum
}

func points(card card.ID) uint8 {
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

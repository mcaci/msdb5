package briscola

import (
	"github.com/mcaci/ita-cards/set"
)

// Sum func
func Score(hand set.Cards) (sum uint8) {
	for _, c := range hand {
		sum += Points(c)
	}
	return
}

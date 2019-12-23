package score

import (
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/briscola"
)

// Sum func
func Sum(hand set.Cards) (sum uint8) {
	for _, c := range hand {
		sum += briscola.Points(c)
	}
	return
}

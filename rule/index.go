package rule

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
)

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable set.Cards, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if DoesOtherCardWin(base, other, briscola) {
			base = other
			max = i
		}
	}
	return uint8(max)
}

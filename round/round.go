package round

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable card.Cards, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if rule.DoesOtherCardWin(base, other, briscola) {
			base = other
			max = i
		}
	}
	return uint8(max)
}

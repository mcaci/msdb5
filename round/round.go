package round

import "github.com/nikiforosFreespirit/msdb5/card"

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable card.Cards, briscola card.Seed) uint8 {
	baseID := cardsOnTheTable[0]
	max := 0
	for i, otherID := range cardsOnTheTable {
		base, _ := card.ByID(baseID)
		other, _ := card.ByID(otherID)
		if DoesOtherCardWin(base, other, briscola) {
			baseID = otherID
			max = i
		}
	}
	return uint8(max)
}

package round

import "github.com/nikiforosFreespirit/msdb5/card"

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable [5]*card.Card, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		onlyOtherCardIsBriscola := !base.IsBriscola(briscola) && other.IsBriscola(briscola)
		otherCardIsBiggerAndOfSameSeed := card.DoesOtherCardWin(base, other)
		if onlyOtherCardIsBriscola || otherCardIsBiggerAndOfSameSeed {
			base = other
			max = i
		}
	}
	return uint8(max)
}

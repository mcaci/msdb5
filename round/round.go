package round

import "github.com/nikiforosFreespirit/msdb5/card"

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable [5]*card.Card, briscola card.Seed) uint8 {
	base := cardsOnTheTable[0]
	max := 0
	for i, other := range cardsOnTheTable {
		if doesOtherCardWin(base, other, briscola) {
			base = other
			max = i
		}
	}
	return uint8(max)
}

func doesOtherCardWin(base, other *card.Card, briscola card.Seed) bool {
	onlyOtherCardIsBriscola := !base.IsBriscola(briscola) && other.IsBriscola(briscola)
	otherCardIsBiggerAndOfSameSeed := DoesOtherCardWin(base, other)
	return onlyOtherCardIsBriscola || otherCardIsBiggerAndOfSameSeed
}

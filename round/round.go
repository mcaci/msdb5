package round

import "github.com/nikiforosFreespirit/msdb5/card"

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable [5]*card.Card, briscola card.Seed) uint8 {
	winningCard := cardsOnTheTable[0]
	max := 0
	for i, card := range cardsOnTheTable {
		onlyOtherCardIsBriscola := !winningCard.IsBriscola(briscola) && card.IsBriscola(briscola)
		otherCardIsBiggerAndOfSameSeed := winningCard.Compare(*card) < 0
		if onlyOtherCardIsBriscola || otherCardIsBiggerAndOfSameSeed {
			winningCard = card
			max = i
		}
	}
	return uint8(max)
}

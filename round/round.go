package round

import "github.com/nikiforosFreespirit/msdb5/card"

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable [5]*card.Card, briscola card.Seed) uint8 {
	winningCard := *cardsOnTheTable[0]
	max := 0
	for i, card := range cardsOnTheTable {
		if !winningCard.IsBriscola(briscola) && card.IsBriscola(briscola) {
			winningCard = *card
			max = i
		} else if winningCard.Compare(*card) < 0 {
			winningCard = *card
			max = i
		} else {
		}
	}
	return uint8(max)
}

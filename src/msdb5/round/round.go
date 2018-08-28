package round

import "msdb5/card"

// IndexOfWinningCard func
func IndexOfWinningCard(cardsOnTheTable [5]*card.Card, briscola card.Seed) uint8 {
	return maxCardIndex(cardsOnTheTable)
}

func maxCardIndex(cards [5]*card.Card) uint8 {
	maxCard := *cards[0]
	max := 0
	for i, card := range cards {
		if maxCard.Compare(*card) < 0 {
			maxCard = *card
			max = i
		}
	}
	return uint8(max)
}

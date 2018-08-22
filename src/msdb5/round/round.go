package round

import "msdb5/card"

func declareWinner(first, second, third, fourth, fifth *card.Card, briscola card.Seed) uint8 {
	cards := []*card.Card{first, second, third, fourth, fifth}
	return maxCardIndex(cards)
}

func maxCardIndex(cards []*card.Card) uint8 {
	maxCard := card.Card{}
	max := 0
	for i, card := range cards {
		if card.Compare(maxCard) > 0 {
			maxCard = *card
			max = i
		}
	}
	return uint8(max)
}

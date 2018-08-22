package round

import "msdb5/card"

func declareWinner(first, second, third, fourth, fifth *card.Card, briscola card.Seed) uint8 {
	cards := []*card.Card{first, second, third, fourth, fifth}
	return maxCardIndex(cards)
}

func maxCardIndex(cards []*card.Card) uint8 {
	for i, card := range cards {
		
	}
	return 0
}

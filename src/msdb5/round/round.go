package round

import "msdb5/card"

func declareWinner(first, second, third, fourth, fifth *card.Card, briscola card.Seed) uint8 {
	winnerCode := -1
	firstCode := 0
	secondCode := 1

	if first.Compare(*second) > 0 {
		winnerCode = firstCode
	} else {
		winnerCode = secondCode
	}
	return uint8(winnerCode)
}

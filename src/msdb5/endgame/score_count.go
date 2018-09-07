package endgame

import "msdb5/card"

// CountPoints func
func CountPoints(cards ...*card.Card) uint8 {
	var sum uint8
	for _, card := range cards {
		sum += card.Points()
	}
	return sum
}

package endgame

import "msdb5/card"

// CountPoints func
func CountPoints(c ...*card.Card) uint8 {
	if len(c) > 0 {
		return c[0].Points()
	}
	return 0
}

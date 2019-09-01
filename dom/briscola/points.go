package briscola

import "github.com/mcaci/ita-cards/card"

// Points func
func Points(id card.Item) uint8 {
	switch id.Number() {
	case 1:
		return 11
	case 3:
		return 10
	case 8:
		return 2
	case 9:
		return 3
	case 10:
		return 4
	default:
		return 0
	}
}

package briscola

import "github.com/mcaci/ita-cards/card"

// Points func
func Points(id card.Item) uint8 {
	var points = map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}
	return points[id.Number()]
}

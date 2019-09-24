package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

// Points func
func Points(id card.Item) uint8 {
	var points = map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}
	return points[id.Number()]
}

// Serie func
func Serie(briscola card.Seed) set.Cards {
	serie := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	cards := make(set.Cards, len(serie))
	for i, id := range serie {
		cards[i] = *card.MustID(id + 10*uint8(briscola))
	}
	return cards
}

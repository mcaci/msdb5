package briscola

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

// Card represents the Briscola Card
type Card struct{ card.Item }

// MustID creates a new card Item (a card with number and seed)
// from an id ranging from 1 to 40, panics if outside
func MustID(n uint8) *Card { return &Card{Item: *card.MustID(n)} }

// Points returns the value of each card number according to Briscola rules
func Points(scorer interface{ Number() uint8 }) uint8 {
	var points = map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}
	return points[scorer.Number()]
}

type Pnts uint8

func (p *Pnts) GetPoints() uint32 { return uint32(*p) }

// Score computes the total score for a cardset according to Briscola rules
func Score(cards set.Cards) *Pnts {
	sum := Pnts(0)
	for _, c := range cards {
		sum += Pnts(Points(c))
	}
	return &sum
}

// FinalScore computes the total score for a cardset according to Briscola rules
func FinalScore(scorers []*card.Item) (sum uint8) {
	for _, s := range scorers {
		sum += Points(s)
	}
	return
}

// Serie computes the list of card numbers from higher to lower points value
func Serie(briscola interface{ Seed() card.Seed }) set.Cards {
	serie := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	cards := make(set.Cards, len(serie))
	for i, id := range serie {
		cards[i] = *card.MustID(id + 10*uint8(briscola.Seed()))
	}
	return cards
}

package briscola_test

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v3/briscola"
)

// MustID creates a new card Item (a card with number and seed)
// from an id ranging from 1 to 40, panics if outside
func mustID(n uint8) *briscola.Card { return &briscola.Card{Item: *card.MustID(n)} }

func TestWinnerIndex(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		cards    set.Cards
		briscola briscola.Card
		winID    uint8
	}{
		"first card wins":                        {cards: *set.NewMust(1, 2), briscola: *mustID(1)},
		"other card wins":                        {cards: *set.NewMust(2, 1), briscola: *mustID(1), winID: 1},
		"first card wins as other has diff seed": {cards: *set.NewMust(22, 11), briscola: *mustID(31)},
		"other card wins as briscola":            {cards: *set.NewMust(21, 12), briscola: *mustID(11), winID: 1},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			cards := make(set.Cards, len(tc.cards))
			for i := range cards {
				cards[i] = tc.cards[i]
			}
			actual := briscola.Winner(cards, tc.briscola.Seed())
			if tc.winID != actual {
				t.Errorf("Expecting player %d to win but %d won instead. Input: %v", tc.winID, actual, tc)
			}
		})
	}
}

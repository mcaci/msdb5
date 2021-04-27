package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

var wintestcases = map[string]struct {
	cards    set.Cards
	briscola Card
	winID    uint8
}{
	"first card wins":                        {cards: *set.NewMust(1, 2), briscola: *MustID(1)},
	"other card wins":                        {cards: *set.NewMust(2, 1), briscola: *MustID(1), winID: 1},
	"first card wins as other has diff seed": {cards: *set.NewMust(22, 11), briscola: *MustID(31)},
	"other card wins as briscola":            {cards: *set.NewMust(21, 12), briscola: *MustID(11), winID: 1},
}

func TestWinCondition(t *testing.T) {
	for name, tc := range wintestcases {
		t.Run(name, func(t *testing.T) {
			actual := IndexOfWinningCard(tc.cards, tc.briscola.Seed())
			if tc.winID != actual {
				t.Errorf("Expecting player %d to win but %d won instead. Input: %v", tc.winID, actual, tc)
			}
		})
	}
}

func TestWinnerIndex(t *testing.T) {
	for name, tc := range wintestcases {
		t.Run(name, func(t *testing.T) {
			cards := make([]*card.Item, len(tc.cards))
			for i := range cards {
				cards[i] = &tc.briscola.Item
			}
			actual := Winner(cards, tc.briscola.Seed())
			if tc.winID != actual {
				t.Errorf("Expecting player %d to win but %d won instead. Input: %v", tc.winID, actual, tc)
			}
		})
	}
}

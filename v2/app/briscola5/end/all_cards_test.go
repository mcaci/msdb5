package end

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola/player"
	"github.com/mcaci/msdb5/v2/dom/briscola/team"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func TestCollect(t *testing.T) {
	p1 := player.New(&player.Options{For2P: true})
	p1.Hand().Add(*card.MustID(11))
	testcases := map[string]struct {
		cardSetter interface{ Pile() *set.Cards }
		expected   *set.Cards
	}{
		"Test with AllCards": {
			cardSetter: newAllCards(team.Players{p1}, briscola5.Side{Cards: *set.NewMust(5)}, newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 6))),
			expected:   set.NewMust(1, 2, 3, 4, 5, 6, 11),
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, expected := *tc.cardSetter.Pile(), *tc.expected
			if len(actual) != len(expected) {
				t.Errorf("expecting same length but found: actual (%d), expected (%d).", len(actual), len(expected))
			}
		nextItem:
			for i := range actual {
				var found bool
				for j := range expected {
					if actual[i] != expected[j] {
						continue
					}
					found = true
					continue nextItem
				}
				if !found {
					t.Errorf("expecting to find item %v in (%v) but was not found: actual slice is (%d).", actual[i], expected, actual)
				}
			}
		})
	}
}

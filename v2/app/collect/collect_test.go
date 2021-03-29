package collect

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func TestCollect(t *testing.T) {
	testcases := map[string]struct {
		cardSetter interface{ Set() *set.Cards }
		expected   *set.Cards
	}{
		"Test with empty RoundCards": {
			cardSetter: NewRoundCards(&set.Cards{}),
			expected:   &set.Cards{},
		},
		"Test with 5 RoundCards": {
			cardSetter: NewRoundCards(&set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)}),
			expected:   &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
		},
		"Test with AllCards": {
			cardSetter: NewAllCards(team.Players{}, &set.Cards{}, &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)}),
			expected:   &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			actual, expected := *tc.cardSetter.Set(), *tc.expected
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

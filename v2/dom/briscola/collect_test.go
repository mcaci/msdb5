package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

type testPile set.Cards

func (p *testPile) Pile() *set.Cards { return (*set.Cards)(p) }

func TestCollect(t *testing.T) {
	testcases := map[string]testPile{
		"Test with 1 card":  {*card.MustID(1)},
		"Test with 5 cards": {*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
	}
	for name, tp := range testcases {
		t.Run(name, func(t *testing.T) {
			expected, actual := tp, testPile{}
			Collect(&tp, &actual)
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
					t.Errorf("expecting to find item %v in (%v) but was not found: actual slice is (%v).", actual[i], expected, actual)
				}
			}
		})
	}
}

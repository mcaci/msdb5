package exchange

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
)

func TestExchangeRound(t *testing.T) {
	testcases := map[string]struct {
		Hand, Side *set.Cards
		hIdx, sIdx int
	}{
		"Test with len 1": {Hand: set.NewMust(1), Side: set.NewMust(2)},
		"Test with len 2": {Hand: set.NewMust(1, 3), Side: set.NewMust(2, 4)},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			bCard, aCard := (*tc.Hand)[tc.hIdx], (*tc.Side)[tc.sIdx]
			Round(tc)
			if (*tc.Hand)[tc.hIdx] != aCard {
				t.Errorf("Expecting exchange to have happened, but not: %v %v %v %v", *tc.Hand, *tc.Side, bCard, aCard)
			}
			if (*tc.Side)[tc.sIdx] != bCard {
				t.Errorf("Expecting exchange to have happened, but not: %v %v %v %v", *tc.Hand, *tc.Side, bCard, aCard)
			}
		})
	}
}

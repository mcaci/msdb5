package exchange

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func TestExchangeRound(t *testing.T) {
	testcases := map[string]struct {
		start struct {
			Hand, Side *set.Cards
			hIdx, sIdx int
		}
		aHand, aSide *set.Cards
	}{
		"Test with len 1": {
			start: struct {
				Hand *set.Cards
				Side *set.Cards
				hIdx int
				sIdx int
			}{Hand: &set.Cards{*card.MustID(1)}, Side: &set.Cards{*card.MustID(2)}},
			aHand: &set.Cards{*card.MustID(2)}, aSide: &set.Cards{*card.MustID(1)},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			Round(tc.start)
			if (*tc.start.Hand)[tc.start.hIdx] != (*tc.aSide)[tc.start.hIdx] {
				t.Errorf("Expecting exchange to have happened, but not: %v %v %v %v", *tc.start.Hand, *tc.start.Side, tc.aHand, tc.aSide)
			}
		})
	}

}

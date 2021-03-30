package play

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type opts struct {
	hands [5]set.Cards
}

func testplayers(opt *opts) team.Players {
	pls := make(team.Players, 5)
	for i := range pls {
		pls[i] = player.New()
		pls[i].Hand().Add(opt.hands[i]...)
	}
	return pls
}

type outParams struct {
	OnBoard set.Cards
}

func TestPlayRound(t *testing.T) {
	// playersWithinLimits := testplayers(&opts{[5]set.Cards{{*card.MustID(1)}, {}, {}, {}, {}}})
	// playersWithinLimitsAndSpreadCards := testplayers(&opts{[5]set.Cards{{*card.MustID(1), *card.MustID(2)}, {*card.MustID(3)}, {}, {}, {}}})
	// playersBeyondLimits := testplayers(&opts{[5]set.Cards{{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4)}, {}, {}, {}, {}}})
	testcases := map[string]struct {
		in  RoundOptions
		out outParams
	}{
		"Test all players with empty hands": {in: RoundOptions{}, out: outParams{}},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			out := Round(&tc.in)
			if len(tc.out.OnBoard) != len(out.OnBoard) {
				t.Errorf("Expected and actual auction results didn't match: Expected (%v), Actual (%v). Input (%v)", tc.out, out, tc.in)
			}
		})
	}
}

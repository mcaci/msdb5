package companion

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type opts struct {
	names [5]string
	hands [5]set.Cards
}

func testplayers(opt *opts) team.Players {
	pls := make(team.Players, 5)
	for i := range pls {
		pls[i] = player.New()
		pls[i].RegisterAs(opt.names[i])
		pls[i].Hand().Add(opt.hands[i]...)
	}
	return pls
}

func TestCompanionRound(t *testing.T) {
	names := [5]string{"a", "b", "c", "d", "e"}
	testcases := map[string]struct {
		crd *card.Item
		pls team.Players
		cmp int
	}{
		"nominal case or choosing self": {
			crd: card.MustID(1),
			pls: testplayers(&opts{
				names: names,
				hands: [5]set.Cards{{}, {*card.MustID(1)}, {}, {}, {}},
			}),
			cmp: 1,
		},
		// e.g. card is in side deck
		// In this case it fallbacks to player 0, for now
		"error case": {
			crd: card.MustID(1),
			pls: testplayers(&opts{names: names}),
			cmp: 0,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			out := Round(tc.crd, tc.pls)
			if tc.crd != out.Briscola {
				t.Error("Unexpected error on the transfer of the briscola card")
			}
			if tc.pls[tc.cmp].Name() != out.Companion.Name() {
				t.Errorf("Unexpected error on the companion selection: expecting (%v), actual (%v)", tc.cmp, out.Companion)
			}
		})
	}
}

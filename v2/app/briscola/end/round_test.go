package end

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type opts struct {
	hands [5]set.Cards
}

type testPlayers team.Players

func testplayers(opt *opts) testPlayers {
	pls := team.New(2)
	for i := range *pls {
		(*pls)[i].Hand().Add(opt.hands[i]...)
	}
	return testPlayers(*pls)
}

func TestEndRound(t *testing.T) {
	testcases := map[string]struct {
		in  Opts
		end bool
	}{
		"Test all players with empty hands": {
			in: Opts{
				Players: team.Players(testplayers(&opts{})),
			},
			end: true},
		"Test one player still with cards": {
			in: Opts{
				Players: team.Players(testplayers(&opts{[5]set.Cards{{*card.MustID(1)}, {}}})),
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			end := Cond(&tc.in)
			if tc.end != end {
				t.Errorf("Expecting end condition to be %t but was not. Input info: %v", tc.end, tc.in)
			}
		})
	}
}

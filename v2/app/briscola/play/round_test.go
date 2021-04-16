package play

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
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

func TestPlayRound(t *testing.T) {
	testcases := map[string]struct {
		in  RoundOpts
		out RoundInfo
	}{
		"Test player with empty hands": {
			in: RoundOpts{
				PlHand:      &set.Cards{},
				PlIdx:       0,
				PlayedCards: &briscola.PlayedCards{Cards: &set.Cards{}},
				NPlayers:    5,
			}, out: RoundInfo{
				OnBoard: &briscola.PlayedCards{Cards: &set.Cards{}},
				NextPl:  1,
			}},
		"Test simple round": {
			in: RoundOpts{
				PlHand:      &set.Cards{*card.MustID(1)},
				PlIdx:       2,
				PlayedCards: &briscola.PlayedCards{Cards: &set.Cards{}},
				NPlayers:    5,
			}, out: RoundInfo{
				OnBoard: &briscola.PlayedCards{Cards: set.NewMust(1)},
				NextPl:  3,
			}},
		"Test last action for round": {
			in: RoundOpts{
				PlHand:       &set.Cards{*card.MustID(1), *card.MustID(2)},
				PlIdx:        2,
				CardIdx:      1,
				NPlayers:     5,
				PlayedCards:  &briscola.PlayedCards{Cards: set.NewMust(11, 21, 12, 22)},
				BriscolaCard: *briscola.MustID(23),
			}, out: RoundInfo{
				OnBoard: &briscola.PlayedCards{Cards: set.NewMust(11, 21, 12, 22, 2)},
				NextPl:  4,
				NextRnd: true,
			}},
		"Test self winning round": {
			in: RoundOpts{
				PlHand:       set.NewMust(11, 33, 28),
				PlIdx:        3,
				CardIdx:      0,
				NPlayers:     5,
				PlayedCards:  &briscola.PlayedCards{Cards: set.NewMust(12, 8, 17, 2)},
				BriscolaCard: *briscola.MustID(33),
			}, out: RoundInfo{
				OnBoard: &briscola.PlayedCards{Cards: set.NewMust(12, 8, 17, 2, 11)},
				NextPl:  3,
				NextRnd: true,
			}},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			out := Round(&tc.in)
			if len(*tc.out.OnBoard.Cards) != len(*out.OnBoard.Cards) {
				t.Errorf("OnBoard error: Expected and actual play results didn't match: Expected (%v), Actual (%v). Input (%v)", tc.out, out, tc.in)
			}
			if tc.out.NextPl != out.NextPl {
				t.Errorf("NextPl error: Expected and actual play results didn't match: Expected (%v), Actual (%v). Input (%v)", tc.out, out, tc.in)
			}
			if tc.out.NextRnd != out.NextRnd {
				t.Errorf("NextRnd error: Expected and actual play results didn't match: Expected (%v), Actual (%v). Input (%v)", tc.out, out, tc.in)
			}
		})
	}
}

package end

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type opts struct {
	hands [5]set.Cards
}

type testPlayersWithCallers team.Players

func testplayers(opt *opts) testPlayersWithCallers {
	pls := make(testPlayersWithCallers, 5)
	for i := range pls {
		pls[i] = player.New()
		pls[i].Hand().Add(opt.hands[i]...)
	}
	return pls
}

func (pls testPlayersWithCallers) Caller() *player.Player    { return pls[1] }
func (pls testPlayersWithCallers) Companion() *player.Player { return player.New() }

func TestEndRound(t *testing.T) {
	playersWithinLimits := testplayers(&opts{[5]set.Cards{{*card.MustID(1)}, {}, {}, {}, {}}})
	playersWithinLimitsAndSpreadCards := testplayers(&opts{[5]set.Cards{{*card.MustID(1), *card.MustID(2)}, {*card.MustID(3)}, {}, {}, {}}})
	playersBeyondLimits := testplayers(&opts{[5]set.Cards{{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4)}, {}, {}, {}, {}}})
	testcases := map[string]struct {
		in  Opts
		end bool
	}{
		"Test all players with empty hands": {in: Opts{}, end: true},
		"Test false because round is in progress": {
			in: Opts{Players: team.Players(playersWithinLimits)},
		},
		"Test false because limit not reached yet": {
			in: Opts{
				PlayedCards: set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
				Players:     team.Players(playersBeyondLimits),
			},
		},
		"Test false because no one has briscola cards": {
			in: Opts{
				PlayedCards:  set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
				Players:      team.Players(playersWithinLimits),
				BriscolaCard: card.MustID(11),
			},
		},
		"Test true because one team only has briscola cards": {
			in: Opts{
				PlayedCards:  set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
				Players:      team.Players(playersWithinLimits),
				BriscolaCard: card.MustID(1),
				Callers:      playersWithinLimits,
			},
			end: true,
		},
		"Test false because not only one team only has briscola cards": {
			in: Opts{
				PlayedCards:  set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4), *card.MustID(5)},
				Players:      team.Players(playersWithinLimitsAndSpreadCards),
				BriscolaCard: card.MustID(1),
				Callers:      playersWithinLimitsAndSpreadCards,
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

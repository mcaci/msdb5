package end

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

type opts struct {
	hands [5]set.Cards
}

type testPlayers misc.Players

func testplayers(opt *opts) testPlayers {
	pls := misc.NewPlayers(5)
	for i := range *pls {
		(*pls)[i] = misc.New(&misc.Options{For5P: true})
		(*pls)[i].Hand().Add(opt.hands[i]...)
	}
	return testPlayers(*pls)
}

func (pls *testPlayers) Caller() misc.Player { return (*pls)[1] }
func (pls *testPlayers) Companion() misc.Player {
	return misc.New(&misc.Options{For2P: true})
}

func newPlayedCardsForTest(a *set.Cards) *briscola.PlayedCards {
	b := briscola.NewPlayedCards(5)
	b.Cards = a
	return b
}

func TestEndRound(t *testing.T) {
	playersWithinLimits := testplayers(&opts{[5]set.Cards{{*card.MustID(1)}, {}, {}, {}, {}}})
	playersWithinLimitsAndSpreadCards := testplayers(&opts{[5]set.Cards{{*card.MustID(1), *card.MustID(2)}, {*card.MustID(3)}, {}, {}, {}}})
	playersBeyondLimits := testplayers(&opts{[5]set.Cards{{*card.MustID(1), *card.MustID(2), *card.MustID(3), *card.MustID(4)}, {}, {}, {}, {}}})
	testcases := map[string]struct {
		in  Opts
		end bool
	}{
		"Test all players with empty hands": {
			in: Opts{
				PlayedCards: *newPlayedCardsForTest(&set.Cards{}),
				Players:     misc.Players(testplayers(&opts{})),
			}, end: true},
		"Test false because round is in progress": {
			in: Opts{
				PlayedCards: *newPlayedCardsForTest(&set.Cards{}),
				Players:     misc.Players(playersWithinLimits),
				Callers:     &playersWithinLimits,
			},
		},
		"Test false because limit not reached yet": {
			in: Opts{
				PlayedCards: *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:     misc.Players(playersBeyondLimits),
				Callers:     &playersBeyondLimits,
			},
		},
		"Test false because no one has briscola cards": {
			in: Opts{
				PlayedCards:  *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:      misc.Players(playersWithinLimits),
				Callers:      &playersWithinLimits,
				BriscolaCard: briscola.Card{Item: *card.MustID(11)},
			},
		},
		"Test true because one team only has briscola cards": {
			in: Opts{
				PlayedCards:  *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:      misc.Players(playersWithinLimits),
				Callers:      &playersWithinLimits,
				BriscolaCard: briscola.Card{Item: *card.MustID(1)},
			},
			end: true,
		},
		"Test false because not only one team only has briscola cards": {
			in: Opts{
				PlayedCards:  *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:      misc.Players(playersWithinLimitsAndSpreadCards),
				Callers:      &playersWithinLimitsAndSpreadCards,
				BriscolaCard: briscola.Card{Item: *card.MustID(1)},
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

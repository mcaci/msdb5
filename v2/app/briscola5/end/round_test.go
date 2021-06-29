package end

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
)

type opts struct {
	hands [5]set.Cards
}

type testPlayers briscola5.Players

func testplayers(opt *opts) testPlayers {
	pls := briscola5.NewPlayers()
	for i := range pls.List() {
		pls.At(i).Hand().Add(opt.hands[i]...)
	}
	return testPlayers(*pls)
}

func (pls *testPlayers) Caller() *player.Player { return (*briscola5.Players)(pls).Player(1) }
func (pls *testPlayers) Companion() *player.Player {
	return player.New(&player.Options{For2P: true}).(*player.Player)
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
				Players:     briscola5.Players(testplayers(&opts{})),
			}, end: true},
		"Test false because round is in progress": {
			in: Opts{
				PlayedCards: *newPlayedCardsForTest(&set.Cards{}),
				Players:     briscola5.Players(playersWithinLimits),
			},
		},
		"Test false because limit not reached yet": {
			in: Opts{
				PlayedCards: *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:     briscola5.Players(playersBeyondLimits),
			},
		},
		"Test false because no one has briscola cards": {
			in: Opts{
				PlayedCards:  *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:      briscola5.Players(playersWithinLimits),
				BriscolaCard: briscola.Card{Item: *card.MustID(11)},
			},
		},
		"Test true because one team only has briscola cards": {
			in: Opts{
				PlayedCards:  *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:      briscola5.Players(playersWithinLimits),
				BriscolaCard: briscola.Card{Item: *card.MustID(1)},
			},
			end: true,
		},
		"Test false because not only one team only has briscola cards": {
			in: Opts{
				PlayedCards:  *newPlayedCardsForTest(set.NewMust(1, 2, 3, 4, 5)),
				Players:      briscola5.Players(playersWithinLimitsAndSpreadCards),
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

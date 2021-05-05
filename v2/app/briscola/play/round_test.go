package play

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/pb"
)

func newPlayedCardsForTest(a *set.Cards) *briscola.PlayedCards {
	b := briscola.NewPlayedCards(2)
	b.Cards = a
	return b
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
				PlayedCards: newPlayedCardsForTest(&set.Cards{}),
				NPlayers:    5,
				EndRound:    endDirect,
			}, out: RoundInfo{
				OnBoard: newPlayedCardsForTest(&set.Cards{}),
				NextPl:  1,
			}},
		"Test simple round": {
			in: RoundOpts{
				PlHand:      &set.Cards{*card.MustID(1)},
				PlIdx:       2,
				PlayedCards: newPlayedCardsForTest(&set.Cards{}),
				NPlayers:    5,
				EndRound:    endDirect,
			}, out: RoundInfo{
				OnBoard: newPlayedCardsForTest(set.NewMust(1)),
				NextPl:  3,
			}},
		"Test last action for round": {
			in: RoundOpts{
				PlHand:       &set.Cards{*card.MustID(1), *card.MustID(2)},
				PlIdx:        2,
				CardIdx:      1,
				NPlayers:     5,
				PlayedCards:  newPlayedCardsForTest(set.NewMust(11, 21, 12, 22)),
				BriscolaCard: *briscola.MustID(23),
				EndRound:     endDirect,
			}, out: RoundInfo{
				OnBoard: newPlayedCardsForTest(set.NewMust(11, 21, 12, 22, 2)),
				NextPl:  4,
				NextRnd: true,
			}},
		"Test self winning round": {
			in: RoundOpts{
				PlHand:       set.NewMust(11, 33, 28),
				PlIdx:        3,
				CardIdx:      0,
				NPlayers:     5,
				PlayedCards:  newPlayedCardsForTest(set.NewMust(12, 8, 17, 2)),
				BriscolaCard: *briscola.MustID(33),
				EndRound:     endDirect,
			}, out: RoundInfo{
				OnBoard: newPlayedCardsForTest(set.NewMust(12, 8, 17, 2, 11)),
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

func endDirect(opts *struct {
	PlayedCards  briscola.PlayedCards
	BriscolaCard briscola.Card
}) (*pb.Index, error) {
	pbcards := make(set.Cards, len(*opts.PlayedCards.Cards))
	for i := range pbcards {
		pbcards[i] = (*opts.PlayedCards.Cards)[i]
	}
	return &pb.Index{Id: uint32(briscola.Winner(pbcards, opts.BriscolaCard.Seed()))}, nil
}

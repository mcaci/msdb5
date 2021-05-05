package auction

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

type opts struct {
	folded [5]bool
}

func testplayers(opt *opts) briscola5.Players {
	pls := briscola5.NewPlayers()
	for i := range pls.List() {
		if opt.folded[i] {
			pls.At(i).Fold()
		}
	}
	return *pls
}

type inParams struct {
	curr, prop briscola5.AuctionScore
	currID     uint8
	players    briscola5.Players
	cmpF       func(briscola5.AuctionScore, briscola5.AuctionScore) int8
}

type outParams struct {
	s   briscola5.AuctionScore
	id  uint8
	end bool
}

func TestAuctionRound(t *testing.T) {
	testcases := map[string]struct {
		in  inParams
		out outParams
	}{
		"Nominal case": {
			in: inParams{
				curr:    61,
				prop:    80,
				currID:  0,
				players: testplayers(&opts{}),
				cmpF:    dirCmp,
			},
			out: outParams{
				s:   80,
				id:  1,
				end: false,
			},
		},
		"Player 0 already folded": {
			in: inParams{
				curr:   61,
				prop:   80,
				currID: 0,
				players: testplayers(&opts{
					folded: [5]bool{true, false, false, false, false},
				}),
				cmpF: dirCmp,
			},
			out: outParams{
				s:   61,
				id:  1,
				end: false,
			}},
		"Player 0 gives value below current one": {
			in: inParams{
				curr:    80,
				prop:    74,
				currID:  0,
				players: testplayers(&opts{}),
				cmpF:    dirCmp,
			},
			out: outParams{
				s:   80,
				id:  1,
				end: false,
			},
		},
		"Player 0 is one before last to fold": {
			in: inParams{
				curr:   80,
				prop:   74,
				currID: 0,
				players: testplayers(&opts{
					folded: [5]bool{false, true, false, true, true},
				}),
				cmpF: dirCmp,
			},
			out: outParams{
				s:   80,
				id:  2,
				end: true,
			},
		},
		"Player 3 give value over the max": {
			in: inParams{
				curr:    81,
				prop:    120,
				currID:  3,
				players: testplayers(&opts{}),
				cmpF:    dirCmp,
			},
			out: outParams{
				s:   120,
				id:  3,
				end: true,
			},
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			out := Round(tc.in)
			if out != tc.out {
				t.Errorf("Expected and actual auction results didn't match: Expected (%v), Actual (%v)", tc.out, out)
			}
		})
	}
}

func dirCmp(curr, prop briscola5.AuctionScore) int8 {
	return int8(briscola5.Cmp(curr, prop))
}

package auction

import (
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func Round(r struct {
	curr    auction.Score
	prop    auction.Score
	currID  uint8
	players briscola5.Players
}) struct {
	s   auction.Score
	id  uint8
	end bool
} {
	// Player has folded already, go to next player and exit
	if briscola5.Folded(r.players[r.currID]) {
		return struct {
			s   auction.Score
			id  uint8
			end bool
		}{s: r.curr, id: mustRotateOnNotFolded(r.players, r.currID)}
	}
	var s auction.Score = r.curr
	var id uint8 = r.currID
	var end bool
	switch auction.Cmp(r.curr, r.prop) {
	case auction.GT_ACTUAL:
		// Auction bid is valid: updates score
		s = auction.CmpAndSet(r.curr, r.prop)
		id = mustRotateOnNotFolded(r.players, r.currID)
	case auction.LE_ACTUAL, auction.LT_MIN_SCORE:
		// Player is folded for scoring less or equal than current (or min)
		r.players[r.currID].Fold()
		// End the loop if only one not folded players is left
		id = mustRotateOnNotFolded(r.players, r.currID)
		end = briscola5.Count(r.players, notFolded) == 1
	case auction.GE_MAX_SCORE:
		// Fold everyone if score is 120 or more
		(&othersFold{p: r.players[r.currID], pls: r.players}).Fold()
		s = auction.MAX_SCORE
		end = true
	}
	return struct {
		s   auction.Score
		id  uint8
		end bool
	}{s: s, id: id, end: end}
}

type othersFold struct {
	p   *briscola5.Player
	pls briscola5.Players
}

func (ot *othersFold) Fold() {
	for _, p := range ot.pls {
		if p == ot.p {
			continue
		}
		p.Fold()
	}
}

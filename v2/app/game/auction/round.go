package auction

import (
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Round(curr, prop auction.Score, currID uint8, players team.Players) struct {
	s   auction.Score
	id  uint8
	end bool
} {
	// Player has folded already, go to next player and exit
	if player.Folded(players[currID]) {
		return struct {
			s   auction.Score
			id  uint8
			end bool
		}{s: curr, id: mustRotateOnNotFolded(players, currID)}
	}
	var s auction.Score = curr
	var id uint8 = currID
	var end bool
	switch auction.Cmp(curr, prop) {
	case auction.GT_ACTUAL:
		// Auction bid is valid: updates score
		s = auction.CmpAndSet(curr, prop)
		id = mustRotateOnNotFolded(players, currID)
	case auction.LE_ACTUAL, auction.LT_MIN_SCORE:
		// Player is folded for scoring less or equal than current (or min)
		players[currID].Fold()
		// End the loop if only one not folded players is left
		id = mustRotateOnNotFolded(players, currID)
		end = team.Count(players, notFolded) == 1
	case auction.GE_MAX_SCORE:
		// Fold everyone if score is 120 or more
		(&othersFold{p: players[currID], pls: players}).Fold()
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
	p   *player.Player
	pls team.Players
}

func (ot *othersFold) Fold() {
	for _, p := range ot.pls {
		if p == ot.p {
			continue
		}
		p.Fold()
	}
}

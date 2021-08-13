package auction

import (
	"fmt"
	"log"

	"github.com/mcaci/msdb5/v2/app/misc"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

const (
	MIN int8 = iota - 2
	LE
	GT
	OVER
)

func Round(r struct {
	curr    briscola5.AuctionScore
	prop    briscola5.AuctionScore
	currID  uint8
	players misc.Players
	cmpF    func(briscola5.AuctionScore, briscola5.AuctionScore) int8
}) struct {
	s   briscola5.AuctionScore
	id  uint8
	end bool
} {
	p := r.players[r.currID]
	// Player has folded already, go to next player and exit
	if folded(p) {
		return struct {
			s   briscola5.AuctionScore
			id  uint8
			end bool
		}{s: r.curr, id: mustRotateOnNotFolded(r.players, r.currID)}
	}
	var s briscola5.AuctionScore = r.curr
	var id uint8 = r.currID
	var end bool
	switch r.cmpF(r.curr, r.prop) {
	case GT:
		// Auction bid is valid: updates score
		s = briscola5.CmpAndSet(r.curr, r.prop)
		id = mustRotateOnNotFolded(r.players, r.currID)
	case MIN, LE:
		// Player is folded for scoring less or equal than current (or min)
		p.(*briscola5.Player).Fold()
		// End the loop if only one not folded players is left
		id = mustRotateOnNotFolded(r.players, r.currID)
		end = misc.Count(r.players, notFolded) == 1
	case OVER:
		// Fold everyone if score is 120 or more
		(&othersFold{p: p.(*briscola5.Player), pls: r.players}).Fold()
		s = briscola5.MAX_SCORE
		end = true
	}
	return struct {
		s   briscola5.AuctionScore
		id  uint8
		end bool
	}{s: s, id: id, end: end}
}

type othersFold struct {
	p   misc.Player
	pls misc.Players
}

func (ot *othersFold) Fold() {
	for _, p := range ot.pls {
		if p == ot.p {
			continue
		}
		p.(*briscola5.Player).Fold()
	}
}

func mustRotateOnNotFolded(players misc.Players, from uint8) uint8 {
	id, err := rotateOn(players, from, notFolded)
	if err != nil {
		log.Printf("error found: %v. Exiting.", err)
	}
	return id
}

func rotateOn(players misc.Players, idx uint8, appliesTo func(p misc.Player) bool) (uint8, error) {
	for i := 0; i < 2*len(players); i++ {
		idx = (idx + 1) % uint8(len(players))
		if !appliesTo(players[idx]) {
			continue
		}
		return idx, nil
	}
	return 0, fmt.Errorf("rotated twice on the number of players and no player found in play")
}

func folded(p misc.Player) bool {
	b5p, ok := p.(*briscola5.Player)
	if !ok {
		return false
	}
	return b5p.Folded()
}

func notFolded(p misc.Player) bool {
	b5p, ok := p.(*briscola5.Player)
	if !ok {
		return false

	}
	return !b5p.Folded()
}

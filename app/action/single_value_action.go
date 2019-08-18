package action

import (
	"strconv"

	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

type Auction struct {
	curr   *player.Player
	score  auction.Score
	update func(auction.Score)
}

func (a Auction) ValueSet(val string) {
	score, err := strconv.Atoi(val)
	toFold := player.Folded(a.curr) || err != nil || !auction.CheckScores(a.score, auction.Score(score))
	if toFold {
		a.curr.Fold()
	}
	newScore := auction.Update(a.score, auction.Score(score))
	a.update(newScore)
}

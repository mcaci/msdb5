package phase

import (
	"strconv"

	"github.com/mcaci/msdb5/dom/auction"
)

type auctioner interface {
	Folded() bool
	AuctionScore() *auction.Score
}

func Auction(rq valueProvider, auct auctioner) Data {
	score, err := strconv.Atoi(rq.Value())
	newScore := auction.Update(*auct.AuctionScore(), auction.Score(score))
	return Data{score: newScore, toFold: auct.Folded() || err != nil || !auction.CheckScores(*auct.AuctionScore(), auction.Score(score))}
}

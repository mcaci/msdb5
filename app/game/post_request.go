package game

import (
	"github.com/mcaci/msdb5/dom/auction"
)

func postJoin(nameProvider interface{ Name() string },
	action interface{ RegisterAs(string) }) {
	action.RegisterAs(nameProvider.Name())
}

func postAuctionFold(action interface{ Fold() }) {
	action.Fold()
}

func postAuctionScore(scoreProvider interface{ Score() auction.Score },
	action interface{ SetAuction(auction.Score) }) {
	action.SetAuction(scoreProvider.Score())
}

package game

import (
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

func PostJoin(namer interface{ Name() string }, joiner interface{ CurrentPlayer() *player.Player }) {
	joiner.CurrentPlayer().RegisterAs(namer.Name())
}

func PostAuctionFold(auctioner interface{ CurrentPlayer() *player.Player }) {
	auctioner.CurrentPlayer().Fold()
}

func PostAuctionScore(scoreProvider interface{ Score() auction.Score },
	effector interface{ SetAuction(auction.Score) }) {
	effector.SetAuction(scoreProvider.Score())
}

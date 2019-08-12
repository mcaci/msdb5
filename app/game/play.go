package game

import (
	"fmt"

	"github.com/mcaci/msdb5/app/cardaction"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

type auctionData struct {
	pl    *player.Player
	score *auction.Score
}

func (a auctionData) Folded() bool                 { return player.Folded(a.pl) }
func (a auctionData) AuctionScore() *auction.Score { return a.score }

func (g *Game) play(rq *request.Req) {
	switch g.Phase() {
	case phase.Joining:
		data := phase.Join(rq)
		postJoin(data, g.CurrentPlayer())
	case phase.InsideAuction:
		data := phase.Auction(rq, auctionData{g.CurrentPlayer(), g.AuctionScore()})
		if data.ToFold() {
			postAuctionFold(g.CurrentPlayer())
			return
		}
		postAuctionScore(data, g)
	}
}

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

func (g *Game) playCard(rq *request.Req) error {
	var a cardaction.Actioner
	switch g.Phase() {
	case phase.ExchangingCards:
		a = cardaction.Exch{g.SideDeck()}
	case phase.ChoosingCompanion:
		a = cardaction.Comp{g.SetBriscola, g.SetCompanion}
	case phase.PlayingCards:
		a = cardaction.Play{g.PlayedCards()}
	default:
		return fmt.Errorf("Action %s not valid", rq.Action())
	}
	return cardaction.CardAction(rq, g.Players(), a)
}

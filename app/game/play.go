package game

import (
	"strconv"

	"github.com/mcaci/msdb5/app/cardaction"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

func (g *Game) play(rq *request.Req) {
	switch g.Phase() {
	case phase.Joining:
		g.CurrentPlayer().RegisterAs(rq.Value())
	case phase.InsideAuction:
		score, err := strconv.Atoi(rq.Value())
		toFold := player.Folded(g.CurrentPlayer()) || err != nil || !auction.CheckScores(*g.AuctionScore(), auction.Score(score))
		if toFold {
			g.CurrentPlayer().Fold()
			return
		}
		newScore := auction.Update(*g.AuctionScore(), auction.Score(score))
		g.SetAuction(newScore)
	}
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
	}
	return cardaction.CardAction(rq, g.Players(), a)
}

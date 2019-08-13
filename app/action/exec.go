package action

import (
	"strconv"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
)

// Exec func
func Exec(g actor, rq cardValueProvider) error {
	switch g.Phase() {
	case phase.Joining:
		g.CurrentPlayer().RegisterAs(rq.Value())
		return nil
	case phase.InsideAuction:
		score, err := strconv.Atoi(rq.Value())
		toFold := player.Folded(g.CurrentPlayer()) || err != nil || !auction.CheckScores(*g.AuctionScore(), auction.Score(score))
		if toFold {
			g.CurrentPlayer().Fold()
			return nil
		}
		newScore := auction.Update(*g.AuctionScore(), auction.Score(score))
		g.SetAuction(newScore)
		return nil
	}
	var a Actioner
	switch g.Phase() {
	case phase.ExchangingCards:
		a = Exch{g.SideDeck()}
	case phase.ChoosingCompanion:
		a = Comp{g.SetBriscola, g.SetCompanion}
	case phase.PlayingCards:
		a = Play{g.PlayedCards()}
	default:
		return nil
	}
	return CardAction(rq, g.Players(), a)
}

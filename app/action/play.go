package action

import (
	"github.com/mcaci/msdb5/app/phase"
)

// Play func
func Play(g actor, rq cardValueProvider) error {
	switch g.Phase() {
	case phase.Joining:
		SingleValueAction(rq, Join(g.CurrentPlayer().RegisterAs))
		return nil
	case phase.InsideAuction:
		SingleValueAction(rq, Auction{g.CurrentPlayer(), *g.AuctionScore(), g.SetAuction})
		return nil
	}
	var a actioner
	switch g.Phase() {
	case phase.ExchangingCards:
		a = Exch{g.SideDeck()}
	case phase.ChoosingCompanion:
		a = Comp{g.SetBriscola, g.SetCompanion}
	case phase.PlayingCards:
		a = PlayCard{g.PlayedCards()}
	}
	return CardAction(rq, g.Players(), a)
}

package game

import (
	"fmt"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
)

func (g *Game) play(rq *request.Req) error {
	switch g.Phase() {
	case phase.Joining:
		data := phase.Join(rq)
		PostJoin(data, g.CurrentPlayer())
	case phase.InsideAuction:
		data := phase.Auction(rq, auctionData{g.CurrentPlayer(), g.AuctionScore()})
		if data.ToFold() {
			PostAuctionFold(g.CurrentPlayer())
			return nil
		}
		PostAuctionScore(data, g)
	case phase.ExchangingCards:
		if rq.Value() == "0" {
			return nil
		}
		data := phase.CardAction(rq, g.Players())
		plHand := g.players[data.Index()].Hand()
		idx := plHand.Find(data.Card())
		PostExchange(plHand, g.SideDeck(), idx, 0)
	case phase.ChoosingCompanion:
		data := phase.CardAction(rq, g.Players())
		if err := data.CardNotFound(); err != nil {
			return err
		}
		PostCompanionCard(data, g)
		PostCompanionPlayer(data, g)
	case phase.PlayingCards:
		data := phase.CardAction(rq, g.Players())
		plHand := g.players[data.Index()].Hand()
		idx := plHand.Find(data.Card())
		PostCardPlay(plHand, g.PlayedCards(), idx)
	default:
		return fmt.Errorf("Action %s not valid", rq.Action())
	}
	return nil
}

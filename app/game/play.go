package game

import (
	"errors"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
)

func (g *Game) play(rq *request.Req) error {
	switch g.Phase() {
	case phase.Joining:
		data := phase.Join(rq)
		postJoin(data, g.CurrentPlayer())
	case phase.InsideAuction:
		data := phase.Auction(rq, auctionData{g.CurrentPlayer(), g.AuctionScore()})
		if data.ToFold() {
			postAuctionFold(g.CurrentPlayer())
			return nil
		}
		postAuctionScore(data, g)
	case phase.ExchangingCards:
		if rq.Value() == "0" {
			return nil
		}
		data := phase.CardAction(rq, g.Players())
		if err := data.CardErr(); err != nil {
			return err
		}
		postExchangeCard(data, g)
	case phase.ChoosingCompanion:
		if rq.Value() == "0" {
			return errors.New("Value 0 for card allowed only for ExchangingCard phase")
		}
		data := phase.CardAction(rq, g.Players())
		if err := data.CardErr(); err != nil {
			return err
		}
		postCompanionCard(data, g)
		postCompanionPlayer(data, g)
	case phase.PlayingCards:
		if rq.Value() == "0" {
			return errors.New("Value 0 for card allowed only for ExchangingCard phase")
		}
		data := phase.CardAction(rq, g.Players())
		if err := data.CardErr(); err != nil {
			return err
		}
		postCardPlay(data, g)
	}
	return nil
}

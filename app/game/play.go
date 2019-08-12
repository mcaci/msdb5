package game

import (
	"errors"
	"fmt"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
)

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

func (g *Game) playCard(rq *request.Req) error {
	switch g.Phase() {
	case phase.ExchangingCards:
		if rq.Value() == "0" {
			return nil
		}
		data := phase.CardAction(rq, g.Players())
		if err := data.CardErr(); err != nil {
			return err
		}
		postExchange(data, g)
	case phase.ChoosingCompanion:
		if rq.Value() == "0" {
			return errors.New("Value 0 for card allowed only for ExchangingCard phase")
		}
		data := phase.CardAction(rq, g.Players())
		if err := data.CardErr(); err != nil {
			return err
		}
		postCompanion(data, g)
	case phase.PlayingCards:
		if rq.Value() == "0" {
			return errors.New("Value 0 for card allowed only for ExchangingCard phase")
		}
		data := phase.CardAction(rq, g.Players())
		if err := data.CardErr(); err != nil {
			return err
		}
		postPlay(data, g)
	default:
		return fmt.Errorf("Action %s not valid", rq.Action())
	}
	return nil
}

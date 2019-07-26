package game

import (
	"fmt"

	"github.com/mcaci/msdb5/app/request"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/deck"
)

func (g *Game) play(rq *request.Req) error {
	switch g.Phase() {
	case phase.Joining:
		data := phase.Join(rq)
		PostJoin(data, g)
	case phase.InsideAuction:
		data := phase.Auction(rq, auctionData{g.CurrentPlayer(), g.AuctionScore()})
		if data.ToFold() {
			PostAuctionFold(g)
			return nil
		}
		PostAuctionScore(data, g)
	case phase.ExchangingCards:
		if rq.Value() == "0" {
			return nil
		}
		data := phase.Companion(rq, g.Players())
		plHand := g.players[data.PlIdx()].Hand()
		idx := plHand.Find(data.Card())
		func(cards, to *deck.Cards, index, toIndex int) {
			(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
		}(plHand, g.SideDeck(), idx, 0)
	case phase.ChoosingCompanion:
		data := phase.Companion(rq, g.Players())
		if err := data.CardNotFound(); err != nil {
			return err
		}
		PostCompanionCard(data, g)
		PostCompanionPlayer(data, g)
	case phase.PlayingCards:
		data := phase.Companion(rq, g.Players())
		plHand := g.players[data.PlIdx()].Hand()
		idx := plHand.Find(data.Card())
		func(cards, to *deck.Cards, index int) {
			to.Add((*cards)[index])
			*cards = append((*cards)[:index], (*cards)[index+1:]...)
		}(plHand, g.PlayedCards(), idx)
	default:
		return msg.Error(fmt.Sprintf("Action %s not valid", rq.Action()), g.Lang())
	}
	return nil
}

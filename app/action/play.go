package action

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type actor interface {
	AuctionScore() *auction.Score
	CurrentPlayer() *player.Player
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	SideDeck() *set.Cards
	SetAuction(auction.Score)
	SetBriscola(*card.Item)
	SetCompanion(*player.Player)
}

// Play func
func Play(g actor, rq cardValueProvider) error {
	var err error
	switch g.Phase() {
	case phase.Joining:
		singleValueAction(rq, joinData{g.CurrentPlayer()})
	case phase.InsideAuction:
		singleValueAction(rq, auctionData{g.CurrentPlayer(), *g.AuctionScore(), g.SetAuction})
	case phase.ExchangingCards:
		err = cardAction(rq, exchangeData{g.SideDeck(), g.Players()})
	case phase.ChoosingCompanion:
		err = cardAction(rq, companionData{g.SetBriscola, g.SetCompanion, g.Players()})
	case phase.PlayingCards:
		err = cardAction(rq, playCardData{g.PlayedCards(), g.Players()})
	}
	return err
}

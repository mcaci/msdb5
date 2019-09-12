package action

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type gamePlayer interface {
	AuctionScore() *auction.Score
	CurrentPlayer() *player.Player
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	SideDeck() *set.Cards
	SetAuction(auction.Score)
	SetBriscola(*card.Item)
	SetCaller(player.Predicate)
	SetCompanion(*player.Player)
	SetShowSide(bool, uint8)
	Card() (*card.Item, error)
	Value() string
}

// Play func
func Play(g gamePlayer) error {
	var err error
	switch g.Phase() {
	case phase.Joining:
		singleValueAction(g, joinData{g.CurrentPlayer()})
	case phase.InsideAuction:
		singleValueAction(g, auctionData{g.CurrentPlayer(), g.Players(), g.AuctionScore(), g.SetAuction, g.SideDeck(), g.SetShowSide, g.SetCaller})
	case phase.ExchangingCards:
		err = cardAction(g, exchangeData{g.SideDeck(), g.Players()})
	case phase.ChoosingCompanion:
		err = cardAction(g, companionData{g.SetBriscola, g.SetCompanion, g.Players()})
	case phase.PlayingCards:
		err = cardAction(g, playCardData{g.PlayedCards(), g.Players()})
	}
	return err
}

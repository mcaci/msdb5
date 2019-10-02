package action

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type gamePlayer interface {
	CurrentPlayer() *player.Player
	Players() team.Players

	AuctionScore() *auction.Score
	PlayedCards() *set.Cards
	Phase() phase.ID
	SideDeck() *set.Cards
	SetAuction(auction.Score)
	SetBriscola(*card.Item)
	SetCaller(player.Predicate)
	SetCompanion(*player.Player)
	SetShowSide(uint8)
	Card() (*card.Item, error)
	Value() string
}

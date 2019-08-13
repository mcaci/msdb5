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

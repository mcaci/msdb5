package game

import (
	"container/list"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type roundInformer interface {
	AuctionScore() *auction.Score
	Caller() *player.Player
	Companion() *player.Player
	CurrentPlayer() *player.Player
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item
	LastPlaying() *list.List
	IsSideUsed() bool
	SideDeck() *set.Cards
	IsRoundOngoing() bool
}

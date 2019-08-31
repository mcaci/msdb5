package msg

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
	LastPlayer() *player.Player
	Players() team.Players
	PlayedCard() card.Item
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item
	LastPlaying() *list.List

	RoundError() error
	IsSideUsed() bool
	IsSideToShow() bool
	SideDeck() *set.Cards
	SideSubset() *set.Cards
}

type senderInfo struct {
	players team.Players
	origin  string
}

func (s senderInfo) From() string          { return s.origin }
func (s senderInfo) Players() team.Players { return s.players }

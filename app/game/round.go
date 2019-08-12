package game

import (
	"container/list"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
)

type roundInformer interface {
	AuctionScore() *auction.Score
	Caller() *player.Player
	Companion() *player.Player
	CurrentPlayer() *player.Player
	LastPlayer() *player.Player
	Players() team.Players
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item
	Lang() language.Tag
	LastPlaying() *list.List
	IsSideUsed() bool
	SideDeck() *set.Cards
	IsRoundOngoing() bool
}

type requestInformer interface {
	From() string
	Action() string
	Value() string
}

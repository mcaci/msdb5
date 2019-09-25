package msg

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/msg/cons"
	"github.com/mcaci/msdb5/app/msg/ml"
	"github.com/mcaci/msdb5/app/msg/pl"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type roundInformer interface {
	CurrentPlayer() *player.Player
	Players() team.Players

	AuctionScore() *auction.Score
	Caller() *player.Player
	Companion() *player.Player
	LastPlayer() *player.Player
	PlayedCard() card.Item
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item

	RoundError() error
	IsSideUsed() bool
	IsSideToShow() bool
	SideDeck() *set.Cards
	SideSubset() *set.Cards
}

// Notify func
func Notify(g roundInformer, l language.Tag, inputRequest, origin string) {
	go cons.Write(g, inputRequest, origin)
	go ml.Write(g)

	printer := message.NewPrinter(l)
	pl.ToPls(g, printer, inputRequest, origin)
	pl.ToLastPl(g, printer)
	pl.ToNewPl(g, printer)
}

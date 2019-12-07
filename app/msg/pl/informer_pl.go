package pl

import (
	"fmt"
	"io"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/app/msg/score"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

type plInformer interface {
	team.Callers

	CurrentPlayer() *player.Player
	Players() team.Players

	AuctionScore() *auction.Score
	LastPlayer() *player.Player
	PlayedCard() card.Item
	PlayedCards() *set.Cards
	Phase() phase.ID
	Briscola() card.Item

	RoundError() error
	IsSideToShow() bool
	SideDeck() *set.Cards
	SideSubset() *set.Cards
}

func ToPls(g plInformer, printer *message.Printer, inputRequest, origin string) {
	sendToPlayers(g, "-----")

	rErr := g.RoundError()
	if rErr != nil {
		senderPred := player.MatchingHost(origin)
		s := g.Players().At(g.Players().MustFind(senderPred))
		io.WriteString(s, TranslateGameStatus(g, printer))
		io.WriteString(s, translatePlayer(g.CurrentPlayer(), g.Briscola(), printer))
		errMsg := translateErr(g, printer, inputRequest, rErr)
		io.WriteString(s, errMsg)
		return
	}

	if g.IsSideToShow() {
		sideDeckMsg := fmt.Sprintf("%s: %s\n", sideDeckRef(printer), translateCards(*g.SideSubset(), printer))
		sendToPlayers(g, sideDeckMsg)
	}

	// send logs
	gameStatusMsg := TranslateGameStatus(g, printer)
	sendToPlayers(g, gameStatusMsg)

	if g.Phase() != phase.End {
		return
	}

	// process end game
	endMsg := translateTeam(g.CurrentPlayer(), g, printer)
	sendToPlayers(g, endMsg)
	// compute score
	t1, t2 := g.Players().Part(team.IsInCallers(g))
	scoreMsg := fmt.Sprintf("%s: [%s: %d], [%s: %d]", endRef(printer),
		teams(printer, 0), score.Sum(team.CommonPile(t1)),
		teams(printer, 1), score.Sum(team.CommonPile(t2)))
	sendToPlayers(g, scoreMsg)
}

func sendToPlayers(g interface{ Players() team.Players }, msg string) {
	for _, pl := range g.Players() {
		io.WriteString(pl, msg)
	}
}

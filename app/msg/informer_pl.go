package msg

import (
	"fmt"
	"io"

	"github.com/mcaci/msdb5/app/action"
	"github.com/mcaci/msdb5/app/input"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/score"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

func toPls(g roundInformer, printer *message.Printer, inputRequest, origin string) {
	sendToPlayers(g, "-----")

	rErr := g.RoundError()
	if rErr != nil {
		s := sender(senderInfo{g.Players(), origin})
		io.WriteString(s, TranslateGameStatus(g, printer))
		io.WriteString(s, TranslatePlayer(g.CurrentPlayer(), g, printer))
		errMsg := translateErr(g, printer, inputRequest, rErr)
		io.WriteString(s, errMsg)
		return
	}

	if g.IsSideToShow() {
		sideDeckMsg := printer.Sprintf("Side deck section: (%s)\n", TranslateCards(*g.SideSubset(), printer))
		sendToPlayers(g, sideDeckMsg)
	}

	// send logs
	gameStatusMsg := TranslateGameStatus(g, printer)
	sendToPlayers(g, gameStatusMsg)

	if g.Phase() != phase.End {
		return
	}

	// process end game
	endMsg := TranslateTeam(g.CurrentPlayer(), g, printer)
	sendToPlayers(g, endMsg)
	// compute score
	scoreTeam1, scoreTeam2 := score.Calc(g, g.Players(), briscola.Points)
	scoreMsg := printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
	sendToPlayers(g, scoreMsg)
}

func toLastPl(g roundInformer, printer *message.Printer) {
	if g.LastPlayer() != g.CurrentPlayer() {
		io.WriteString(g.LastPlayer(), TranslatePlayer(g.LastPlayer(), g, printer))
	}
}

func toNewPl(g roundInformer, printer *message.Printer) {
	if g.Phase() == phase.ExchangingCards {
		io.WriteString(g.CurrentPlayer(), TranslateSideDeck(g, g.CurrentPlayer(), printer))
	}
	io.WriteString(g.CurrentPlayer(), TranslatePlayer(g.CurrentPlayer(), g, printer))
}

func translateErr(g roundInformer, printer *message.Printer, inputRequest string, rErr error) string {
	errMsg := fmt.Sprintf("Error: %+v\n", rErr)
	if rErr == phase.ErrUnexpectedPhase {
		_, id := phase.ToID(input.Value(inputRequest))
		errMsg = printer.Sprintf("Phase is not %d but %d", id, g.Phase())
	}
	if rErr == action.ErrUnexpectedPlayer {
		errMsg = printer.Sprintf("Expecting player %s to play", g.CurrentPlayer().Name())
	}
	return errMsg
}

func sendToPlayers(g interface{ Players() team.Players }, msg string) {
	for _, pl := range g.Players() {
		io.WriteString(pl, msg)
	}
}

package msg

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/input"

	"github.com/mcaci/msdb5/app/action/end"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/score"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Notify func
func Notify(g roundInformer, l language.Tag, inputRequest, origin string) {
	for _, pl := range g.Players() {
		io.WriteString(pl, "-----")
		io.WriteString(os.Stdout, inputRequest)
	}

	printer := message.NewPrinter(l)
	rErr := g.RoundError()
	if rErr != nil {
		s := senderInfo{g.Players(), origin}
		errMsg := fmt.Sprintf("Error: %+v\n", rErr)
		if rErr == phase.ErrUnexpectedPhase {
			_, id := phase.ToID(input.Value(inputRequest))
			errMsg = printer.Sprintf("Phase is not %d but %d", id, g.Phase())
		}
		if rErr == team.ErrUnexpectedPlayer {
			errMsg = printer.Sprintf("Expecting player %s to play", g.CurrentPlayer().Name())
		}
		io.WriteString(os.Stdout, errMsg)
		sender := team.Sender(s)
		io.WriteString(sender, TranslateGameStatus(g, printer))
		io.WriteString(sender, CreateInGameMsg(g, g.CurrentPlayer(), l))
		io.WriteString(sender, errMsg)
		return
	}

	if g.IsSideToShow() {
		for _, pl := range g.Players() {
			io.WriteString(pl, printer.Sprintf("Side deck section: (%s)\n", TranslateCards(*g.SideSubset(), printer)))
		}
	}

	// send logs
	senderPlayer := team.Sender(senderInfo{g.Players(), origin})
	io.WriteString(os.Stdout, fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", senderPlayer.Name(), inputRequest, senderPlayer, g))
	if g.LastPlayer() != g.CurrentPlayer() {
		io.WriteString(g.LastPlayer(), CreateInGameMsg(g, g.LastPlayer(), l))
	}
	gameStatusMsg := TranslateGameStatus(g, printer)
	for _, pl := range g.Players() {
		io.WriteString(pl, gameStatusMsg)
	}
	io.WriteString(g.CurrentPlayer(), CreateInGameMsg(g, g.CurrentPlayer(), l))

	io.WriteString(HandleMLData(g))

	if g.Phase() != phase.End {
		return
	}

	// process end game
	endMsg := TranslateTeam(end.Player(g), g, printer)
	for _, pl := range g.Players() {
		io.WriteString(pl, endMsg)
	}

	// compute score
	pilers := make([]score.Piler, len(g.Players()))
	for i, p := range g.Players() {
		pilers[i] = p
	}
	scoreTeam1, scoreTeam2 := score.Calc(g.Caller(), g.Companion(), pilers, briscola.Points)
	scoreMsg := printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
	for _, pl := range g.Players() {
		io.WriteString(pl, scoreMsg)
	}

	io.WriteString(HandleMLData(g))
}

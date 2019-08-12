package game

import (
	"fmt"
	"io"
	"os"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/auction"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) {
	printer := message.NewPrinter(g.Lang())
	rq := request.New(inputRequest, origin)
	report := func(err error) {
		io.WriteString(os.Stdout, fmt.Sprintf("New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), *rq, err))
		io.WriteString(sender(g, rq), fmt.Sprintf("Error: %+v\n", err))
	}

	// verify phase step
	err := request.VerifyPhase(g, rq)
	if err != nil {
		ph, _ := phase.ToID(rq)
		report(msg.UnexpectedPhaseErr(ph, g.Phase(), g.Lang()))
		return
	}

	// verify player step
	err = request.VerifyPlayer(g, rq)
	if err != nil {
		report(msg.UnexpectedPlayerErr(g.CurrentPlayer().Name(), g.Lang()))
		return
	}

	// play step
	g.play(rq)
	g.playCard(rq)

	cardN := auction.SideCards(*g.AuctionScore())
	if phase.InsideAuction == g.Phase() && len(*g.SideDeck()) != 0 && cardN > 0 {
		for _, pl := range g.Players() {
			io.WriteString(pl, printer.Sprintf("Side deck section: (%s)\n", msg.TranslateCards((*g.SideDeck())[:cardN], printer)))
		}
	}

	// log action to file for ml
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		io.WriteString(os.Stdout, err.Error())
		return
	}
	defer f.Close()
	// write to file for ml
	switch g.Phase() {
	case phase.ChoosingCompanion:
		io.WriteString(f, fmt.Sprintf("%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore())))
	case phase.PlayingCards:
		lastPlayed := g.playedCards[len(g.playedCards)-1]
		io.WriteString(f, fmt.Sprintf("%s, %d\n", g.CurrentPlayer().Name(), lastPlayed))
	}

	// end round: next player
	plIndex := nextPlayer(g, rq)
	// next phase
	setCaller := func(p *player.Player) { g.caller = p }
	ph := nextPhase(g, rq, setCaller)
	// clean up
	cleanUp(g, plIndex)
	g.phase = ph
	track.Player(g.LastPlaying(), g.Players()[plIndex])

	// log action to console
	io.WriteString(os.Stdout, fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), *rq, sender(g, rq), g))
	for _, pl := range g.Players() {
		io.WriteString(pl, "-----")
	}
	io.WriteString(g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer()))
	for _, pl := range g.Players() {
		io.WriteString(pl, msg.TranslateGameStatus(g, printer))
	}
	io.WriteString(g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer()))

	if g.phase != phase.End {
		return
	}

	// process end phase
	remainingCards := len(*g.Players()[0].Hand())
	if remainingCards > 0 {
		collect(g)
	}
	// compute score
	pilers := make([]team.Piler, 0)
	for _, p := range g.Players() {
		pilers = append(pilers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), pilers)
	for _, pl := range g.Players() {
		io.WriteString(pl, printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2))
	}
	// write to file
	io.WriteString(f, fmt.Sprintf("%s\n", g.CurrentPlayer().Name()))
}

func sender(g *Game, rq requestInformer) *player.Player {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return g.Players()[index]
}

package game

import (
	"fmt"
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

type proc struct {
	rq      *request.Req
	reports []PlMsg
	err     error
}

// Process2 func
func (g *Game) Process2(inputRequest, origin string) []PlMsg {
	printer := message.NewPrinter(g.Lang())
	pr := proc{}

	rq := request.New(inputRequest, origin)
	pr.rq = rq

	// verify phase step
	if pr.err == nil {
		err := request.VerifyPhase(g, rq)
		// err = msg.UnexpectedPhaseErr(phase.MustID(rq), g.Phase(), g.Lang())
		pr.reportErr(g, err)
	}

	// verify player step
	if pr.err == nil {
		err := request.VerifyPlayer(g, rq)
		// err = msg.UnexpectedPlayerErr(g.CurrentPlayer().Name(), g.Lang())
		pr.reportErr(g, err)
	}

	// play step
	if pr.err == nil {
		g.play(rq)
		err := g.playCard(rq)
		pr.reportErr(g, err)
	}

	if pr.err == nil {
		cardN := auction.SideCards(*g.AuctionScore())
		if phase.InsideAuction == g.Phase() && len(*g.SideDeck()) != 0 && cardN > 0 {
			for _, pl := range g.Players() {
				pl, plMsg := pl, printer.Sprintf("Side deck section: (%s)\n", msg.TranslateCards((*g.SideDeck())[:cardN], printer))
				pr.reports = append(pr.reports, PlMsg{pl, plMsg})
			}
		}

		g.handleMLData() // fake placeholder for memory

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
		cons, consMsg := os.Stdout, fmt.Sprintf("New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), *rq, sender(g, rq), g)
		pr.reports = append(pr.reports, PlMsg{cons, consMsg})
		for _, pl := range g.Players() {
			pl, plMsg := pl, "-----"
			pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		}
		pl, plMsg := g.LastPlayer(), msg.CreateInGameMsg(g, g.LastPlayer())
		pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		for _, pl := range g.Players() {
			pl, plMsg := pl, msg.TranslateGameStatus(g, printer)
			pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		}
		pl, plMsg = g.CurrentPlayer(), msg.CreateInGameMsg(g, g.CurrentPlayer())
		pr.reports = append(pr.reports, PlMsg{pl, plMsg})

		if g.phase != phase.End {
			return pr.reports
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
			pl, plMsg := pl, printer.Sprintf("The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
			pr.reports = append(pr.reports, PlMsg{pl, plMsg})
		}
	}
	return pr.reports
}

func (pr *proc) reportErr(g interface{ Players() team.Players }, err error) {
	cons, consMsg := os.Stdout, fmt.Sprintf("New Action by %s: %s\nError raised: %+v\n", sender(g, pr.rq).Name(), *pr.rq, err)
	pl, plMsg := sender(g, pr.rq), fmt.Sprintf("Error: %+v\n", err)
	pr.reports = append(pr.reports, PlMsg{cons, consMsg})
	pr.reports = append(pr.reports, PlMsg{pl, plMsg})
	pr.err = err
}

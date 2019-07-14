package game

import (
	"fmt"
	"os"

	"github.com/mcaci/msdb5/app/msg"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/play"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) {
	printer := message.NewPrinter(g.Lang())
	rq := request.New(inputRequest, origin)

	// verify phase step
	err := request.VerifyPhase(g, rq)
	if err != nil {
		fmt.Fprintf(os.Stdout, "New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), rq.Action(), err)
		printer.Fprintf(sender(g, rq), "Error: %+v\n", err)
		return
	}

	// verify player step
	err = request.VerifyPlayer(g, rq)
	if err != nil {
		fmt.Fprintf(os.Stdout, "New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), rq.Action(), err)
		printer.Fprintf(sender(g, rq), "Error: %+v\n", err)
		return
	}

	// play step
	setCompanion := func(p *player.Player) { g.companion = p }
	setBriscolaCard := func(c card.ID) { g.briscolaCard = c }
	err = play.Request(g, rq, setCompanion, setBriscolaCard)
	if err != nil {
		fmt.Fprintf(os.Stdout, "New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), rq.Action(), err)
		printer.Fprintf(sender(g, rq), "Error: %+v\n", err)
		return
	}

	// log action to file
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stdout, "New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), rq.Action(), err)
		return
	}
	defer f.Close()
	// write to file
	canLog, text := msg.CreateMlMsg(g)
	if canLog {
		fmt.Fprintf(f, text)
	}

	// end round
	setCaller := func(p *player.Player) { g.caller = p }
	setPhase := func(p phase.ID) { g.phase = p }
	// actions to do post request
	postRequest(g, rq)
	// next player
	nextPlayer(g, rq)
	// next phase
	nextPhase(g, rq, setCaller, setPhase)
	// clean phase
	cleanPhase(g, rq)

	// log action to console
	fmt.Fprintf(os.Stdout, "New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), rq.Action(), sender(g, rq), g)

	// exit if not end game
	if g.phase == phase.End {
		// compute score
		scorers := make([]team.Scorer, 0)
		for _, p := range g.Players() {
			scorers = append(scorers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.Caller(), g.Companion(), scorers...)
		for _, pl := range g.Players() {
			printer.Fprintf(pl, "The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
		}
		// write to file
		canLog, text = msg.CreateMlMsg(g)
		if canLog {
			fmt.Fprintf(f, text)
		}
	}
}

func sender(g *Game, rq requestInformer) *player.Player {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return g.Players()[index]
}

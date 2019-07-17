package game

import (
	"fmt"
	"os"

	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/app/play"
	"github.com/mcaci/msdb5/app/request"
	"github.com/mcaci/msdb5/app/track"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
	"golang.org/x/text/message"
)

// Process func
func (g *Game) Process(inputRequest, origin string) {
	printer := message.NewPrinter(g.Lang())
	rq := request.New(inputRequest, origin)
	report := func(err error) {
		fmt.Fprintf(os.Stdout, "New Action by %s: %s\nError raised: %+v\n", sender(g, rq).Name(), *rq, err)
		printer.Fprintf(sender(g, rq), "Error: %+v\n", err)
	}

	// verify phase step
	err := request.VerifyPhase(g, rq)
	if err != nil {
		report(err)
		return
	}

	// verify player step
	err = request.VerifyPlayer(g, rq)
	if err != nil {
		report(err)
		return
	}

	// play step
	setCompanion := func(p *player.Player) { g.companion = p }
	setBriscolaCard := func(c card.ID) { g.briscolaCard = c }
	err = play.Request(g, rq, setCompanion, setBriscolaCard)
	if err != nil {
		report(err)
		return
	}

	// log action to file for ml
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}
	defer f.Close()
	// write to file for ml
	switch g.Phase() {
	case phase.ChoosingCompanion:
		fmt.Fprintf(f, "%s, %s, %d\n", g.CurrentPlayer().Name(), g.Companion().Name(), *(g.AuctionScore()))
	case phase.PlayingCards:
		lastPlayed := g.playedCards[len(g.playedCards)-1]
		fmt.Fprintf(f, "%s, %d\n", g.CurrentPlayer().Name(), lastPlayed)
	}
	// log action to console
	fmt.Fprintf(os.Stdout, "New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), *rq, sender(g, rq), g)

	// end round: next player
	nextPlIdx := nextPlayer(g, rq)
	// next phase
	setCaller := func(p *player.Player) { g.caller = p }
	ph := nextPhase(g, rq, setCaller)
	// clean up
	cleanUp(g, rq)
	g.phase = ph
	track.Player(g.LastPlaying(), g.Players()[nextPlIdx])

	// process end phase
	if g.phase == phase.End {
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
			printer.Fprintf(pl, "The end - Callers: %d; Others: %d", scoreTeam1, scoreTeam2)
		}
		// write to file
		fmt.Fprintf(f, "%s\n", g.CurrentPlayer().Name())
	}
}

func sender(g *Game, rq requestInformer) *player.Player {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return g.Players()[index]
}

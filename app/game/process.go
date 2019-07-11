package game

import (
	"fmt"
	"os"

	"github.com/nikiforosFreespirit/msdb5/app/end"
	"github.com/nikiforosFreespirit/msdb5/app/msg"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/play"
	"github.com/nikiforosFreespirit/msdb5/app/request"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
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
	end.Round(g, rq, setCaller, setPhase)

	// log action to console
	fmt.Fprintf(os.Stdout, "New Action by %s: %s\nSender info: %+v\nGame info: %+v\n", sender(g, rq).Name(), rq.Action(), sender(g, rq), g)

	// exit if not end game
	if g.phase == phase.End {
		end.Score(g) // process end game
		// write to file
		canLog, text = msg.CreateMlMsg(g)
		if canLog {
			fmt.Fprintf(f, text)
		}
	}
}

type requestInformer interface {
	From() string
}

func sender(g *Game, rq requestInformer) *player.Player {
	index, _ := g.Players().Find(func(p *player.Player) bool { return p.IsSameHost(rq.From()) })
	return g.Players()[index]
}

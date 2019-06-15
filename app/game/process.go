package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/end"
	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/app/phase"
	"github.com/nikiforosFreespirit/msdb5/app/play"
	"github.com/nikiforosFreespirit/msdb5/app/request"
	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Process func
func (g *Game) Process(inputRequest, origin string) {
	notify := func(p *player.Player, msg string) { p.ReplyWith(msg) }
	rq := request.New(inputRequest, origin)

	// verify phase step
	err := request.VerifyPhase(g, rq, notify)
	if err != nil {
		gamelog.SendErrToSender(err, g, rq, notify)
		return
	}

	// verify player step
	err = request.VerifyPlayer(g, rq, notify)
	if err != nil {
		gamelog.SendErrToSender(err, g, rq, notify)
		return
	}

	// play step
	setCompanion := func(p *player.Player) { g.companion = p }
	setBriscolaCard := func(c card.ID) { g.briscolaCard = c }
	err = play.Request(g, rq, setCompanion, setBriscolaCard, notify)
	if err != nil {
		gamelog.SendErrToSender(err, g, rq, notify)
		return
	}

	// log action to file
	gamelog.ToFile(g)

	// end round
	setCaller := func(p *player.Player) { g.caller = p }
	setPhase := func(p phase.ID) { g.phase = p }
	end.Round(g, rq, setCaller, setPhase, notify)

	// log action to console
	gamelog.ToConsole(g, rq)

	// process end game
	if g.phase == phase.End {
		end.Process(g, notify)
	}
}

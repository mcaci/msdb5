package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Process func
func (g *Game) Process(request, origin string) {
	notify := func(p *player.Player, msg string) { p.ReplyWith(msg) }
	rq := newReq(request, origin)

	// verify phase step
	err := verifyPhase(g, rq, notify)
	if err != nil {
		return
	}

	// verify player step
	err = verifyPlayer(g, rq, notify)
	if err != nil {
		return
	}

	// do step
	err = processRequest(g, rq, notify)
	if err != nil {
		return
	}

	// log action to file
	gamelog.Write(g)

	// next player step
	nextPlayer(g, rq, notify)

	// next phase
	nextPhase(g, rq, notify)

	// log action to players
	notifyPlayer(g, rq, notify)
	notifyAll(g, rq, notify)
	gamelog.ToConsole(g, g.sender(rq.From()), rq.Action(), err)

	// clean phase
	cleanPhase(g, rq, notify)

	// process end game
	endGamePhase(g, rq, notify)
}

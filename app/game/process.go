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
		gamelog.SendErrToSender(err, g, rq, notify)
		return
	}

	// verify player step
	err = verifyPlayer(g, rq, notify)
	if err != nil {
		gamelog.SendErrToSender(err, g, rq, notify)
		return
	}

	// play step
	err = processRequest(g, rq, notify)
	if err != nil {
		gamelog.SendErrToSender(err, g, rq, notify)
		return
	}

	// log action to file
	gamelog.ToFile(g)

	// next player step
	nextPlayer(g, rq, notify)

	// next phase
	nextPhase(g, rq, notify)

	// log action to console
	gamelog.ToConsole(g, rq)

	// clean phase
	cleanPhase(g, rq, notify)

	// process end game
	endGamePhase(g, rq, notify)
}

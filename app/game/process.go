package game

import (
	"github.com/nikiforosFreespirit/msdb5/app/gamelog"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Process func
func (g *Game) Process(request, origin string) {
	notify := func(p *player.Player, msg string) { p.ReplyWith(msg) }

	// verify phase step
	err := verifyPhase(g, request, origin, notify)
	if err != nil {
		return
	}

	// verify player step
	err = verifyPlayer(g, request, origin, notify)
	if err != nil {
		return
	}

	// do step
	err = processRequest(g, request, origin, notify)
	if err != nil {
		return
	}

	// log action to file
	gamelog.Write(g)

	// next player step
	nextPlayer(g, request, origin, notify)

	// next phase
	nextPhase(g, request)

	// log action to players
	notifyPlayer(g, request, origin, notify)
	notifyAll(g, request, origin, notify)
	gamelog.ToConsole(g, g.sender(origin), request, err)

	// clean phase
	cleanPhase(g, request, origin, notify)

	// process end game
	endGamePhase(g, request, origin, notify)
}

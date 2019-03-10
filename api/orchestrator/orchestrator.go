package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
)

// Action interface
func (g *Game) Action(request, origin string) (all []display.Info, me []display.Info, err error) {
	data := strings.Split(string(request), "#")
	playerInTurn := g.playerInTurn
	switch data[0] {
	case "Join":
		err, all, me = g.join(data[0], data[1], origin), g.Info(), g.players[playerInTurn].Info()
	case "Auction":
		err, all, me = g.raiseAuction(data[0], data[1], origin), g.Info(), g.players[playerInTurn].Info()
	case "Companion":
		err, all, me = g.nominate(data[0], data[1], data[2], origin), g.Info(), g.players[playerInTurn].Info()
	case "Card":
		err, all, me = g.play(data[0], data[1], data[2], origin), g.Info(), g.players[playerInTurn].Info()
	}
	logEndRound(g, request, origin, err)
	if g.phase == end {
		all, me, err = endGame(g)
	}
	return
}

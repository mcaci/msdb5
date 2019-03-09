package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
)

// Action interface
func (g *Game) Action(request, origin string) ([]display.Info, []display.Info, error) {
	data := strings.Split(string(request), "#")
	var err error
	switch data[0] {
	case "Join":
		err = g.Join(data[1], origin)
	case "Auction":
		err = g.RaiseAuction(data[1], origin)
	case "Companion":
		err = g.Nominate(data[1], data[2], origin)
	case "Card":
		err = g.Play(data[1], data[2], origin)
	}
	logEndRound(g, request, origin, err)
	infoForAllPlayers := g.Info()
	infoForSinglePlayer := g.players[g.playerInTurn].Info()
	if g.phase == end {
		infoForAllPlayers, infoForSinglePlayer, err = endGame(g)
	}
	return infoForAllPlayers, infoForSinglePlayer, err
}

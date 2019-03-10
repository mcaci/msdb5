package orchestrator

import (
	"errors"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) join(request, origin string) (all []display.Info, me []display.Info, err error) {
	data := strings.Split(request, "#")
	action := data[0]
	name := data[1]
	playerInTurn := g.playerInTurn
	if action == "Join" {
		find := isNameEmpty
		do := func(p *player.Player) error {
			p.Join(name, origin)
			return nil
		}
		nextPlayerSupplier := func() uint8 { return (g.playerInTurn + 1) % 5 }
		nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
		err = g.playPhase(joining, find, do, nextPlayerSupplier, nextPhasePredicate)
		return g.Info(), g.players[playerInTurn].Info(), err
	}
	return g.Info(), g.players[playerInTurn].Info(), errors.New("JOIN action not invoked")
}

package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) join(request, origin string) (all []display.Info, me []display.Info, err error) {
	playerInTurn := g.playerInTurn
	info := g.joinData(request, origin)
	return g.Info(), g.players[playerInTurn].Info(), g.playPhase(info)
}

func (g *Game) joinData(request, origin string) dataPhase {
	data := strings.Split(request, "#")
	name := data[1]
	do := func(p *player.Player) error {
		p.Join(name, origin)
		return nil
	}
	nextPlayerSupplier := func() uint8 { return (g.playerInTurn + 1) % 5 }
	nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
	return dataPhase{joining, isNameEmpty, do, nextPlayerSupplier, nextPhasePredicate}
}

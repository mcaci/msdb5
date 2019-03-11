package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) nominate(request, origin string) (all []display.Info, me []display.Info, err error) {
	_, err = cardAction(request)
	if err != nil {
		return
	}
	playerInTurn := g.playerInTurn
	info := g.nominateData(request, origin)
	return g.Info(), g.players[playerInTurn].Info(), g.playPhase(info)
}

func (g *Game) nominateData(request, origin string) dataPhase {
	c, _ := cardAction(request)
	phase := companionChoice
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	do := func(p *player.Player) (err error) {
		if err != nil {
			return
		}
		return g.setCompanion(c)
	}
	nextPlayerSupplier := func() uint8 { return g.playerInTurn }
	nextPhasePredicate := func() bool { return true }
	return dataPhase{phase, find, do, nextPlayerSupplier, nextPhasePredicate}
}

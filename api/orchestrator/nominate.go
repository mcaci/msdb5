package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
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
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) error { return g.setCompanion(c) }
	nextPlayerOperator := func(playerInTurn uint8) uint8 { return playerInTurn }
	nextPhasePredicate := func() bool { return nominateNextPhase(nil, nil) }
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate}
}

func nominateNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return true
}

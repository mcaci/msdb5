package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

type phaseData struct {
	phase              phase
	find               func(*player.Player) bool
	do                 func(*player.Player) error
	nextPlayerOperator func(uint8) uint8
	nextPhasePredicate func(playerset.Players, func(*player.Player) bool) bool
	playerPredicate    func(*player.Player) bool
}

func (g *Game) playPhase(info phaseData) (err error) {
	if err = g.phaseCheck(info.phase); err != nil {
		return
	}
	p, err := g.players.Find(info.find)
	if err != nil {
		return
	}
	err = info.do(p)
	if err != nil {
		return
	}
	g.nextPlayer(info.nextPlayerOperator)
	g.nextPhase(info.nextPhasePredicate, info.playerPredicate)
	return
}

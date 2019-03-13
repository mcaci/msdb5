package orchestrator

import (
	"errors"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func playPhase(g *game.Game, info action.Data) (err error) {
	if err = phaseCheck(g.CurrentPhase(), info.Phase()); err != nil {
		return
	}
	p, err := g.Players().Find(info.Find())
	if err != nil {
		return
	}
	err = info.Do()(p)
	if err != nil {
		return
	}
	nextPlayer(g, info.NextPlayerOperator())
	nextPhase(g, info.NextPhasePredicate(), info.PlayerPredicate())
	return
}

func phaseCheck(gamePhase, current game.Phase) (err error) {
	if gamePhase != current {
		err = errors.New("Phase is not " + strconv.Itoa(int(current)))
	}
	return
}

func nextPhase(g *game.Game, predicate func(playerset.Players, func(*player.Player) bool) bool, playerPredicate func(*player.Player) bool) {
	if predicate(g.Players(), playerPredicate) {
		g.IncrementPhase()
	}
}

func nextPlayer(g *game.Game, generateIndex func(uint8) uint8) {
	g.UpdatePlayerInTurn(generateIndex)
}

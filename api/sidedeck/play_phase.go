package sidedeck

import (
	"errors"
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/api/action"
	"github.com/nikiforosFreespirit/msdb5/api/game"
)

func playPhase(g *game.Game, info action.Action) (err error) {
	if err = phaseCheck(g.CurrentPhase(), info.Phase()); err != nil {
		return
	}
	p, err := g.Players().Find(info.Find)
	if err != nil {
		return
	}
	err = info.Do(p)
	if err != nil {
		return
	}
	g.NextPlayer(info.NextPlayer)
	g.NextPhase(info.NextPhase(g.Players(), info))
	return
}

func phaseCheck(gamePhase, current game.Phase) (err error) {
	if gamePhase != current {
		err = errors.New("Phase is not " + strconv.Itoa(int(current)))
	}
	return
}

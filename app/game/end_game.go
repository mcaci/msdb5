package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/app/phase"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func endGamePhase(g *Game, rq *req, notify func(*player.Player, string)) error {
	if g.phase == phase.End {
		scorers := make([]team.Scorer, 0)
		for _, p := range g.players {
			scorers = append(scorers, p)
		}
		scoreTeam1, scoreTeam2 := team.Score(g.caller, g.companion, scorers...)
		for _, pl := range g.players {
			notify(pl, fmt.Sprintf("The end - Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2))
		}
	}
	return nil
}

package game

import (
	"fmt"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
	"github.com/nikiforosFreespirit/msdb5/dom/team"
)

func endGamePhase(g *Game, request, origin string, notify func(*player.Player, string)) error {
	scorers := make([]team.Scorer, 0)
	for _, p := range g.players {
		scorers = append(scorers, p)
	}
	scoreTeam1, scoreTeam2 := team.Score(g.caller, g.companion, scorers...)
	notify(g.CurrentPlayer(), fmt.Sprintf("Callers: %+v; Others: %+v", scoreTeam1, scoreTeam2))
	return nil
}

package orchestrator

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/display"
)

func verifyEndGame(g *Game) bool {
	gameHasEnded := true
	for _, pl := range g.players {
		if len(*pl.Hand()) > 0 {
			gameHasEnded = false
		}
	}
	return gameHasEnded
}

func endGame(g *Game) ([]display.Info, []display.Info, error) {
	caller, _ := g.players.Find(notFolded)
	score1 := caller.Count() + g.companion.Ref().Count()
	score2 := uint8(0)
	for _, pl := range g.players {
		if pl != caller && pl != g.companion.Ref() {
			score2 += pl.Count()
		}
	}
	score1info := display.NewInfo("Callers", ":", strconv.Itoa(int(score1)), ";")
	score2info := display.NewInfo("Others", ":", strconv.Itoa(int(score2)), ";")
	return display.Wrap("Final Score", score1info, score2info), nil, nil
}

package orchestrator

import (
	"strconv"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func (g *Game) endGameCondition(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.All(searchCriteria)
}

func (g *Game) endGame() ([]display.Info, []display.Info, error) {
	caller, _ := g.players.Find(func(p *player.Player) bool { return p.NotFolded() })
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

package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
	"github.com/nikiforosFreespirit/msdb5/team"
)

func (g *Game) endGameCondition(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.All(searchCriteria)
}

func (g *Game) end() ([]display.Info, []display.Info, error) {
	caller, _ := g.players.Find(func(p *player.Player) bool { return p.NotFolded() })
	team1, team2 := new(team.BriscolaTeam), new(team.BriscolaTeam)
	team1.Add(caller, g.companion.Ref())
	for _, pl := range g.players {
		if pl != caller && pl != g.companion.Ref() {
			team2.Add(pl)
		}
	}
	return display.Wrap("Final Score", team1.Info("Callers"), team2.Info("Others")), nil, nil
}

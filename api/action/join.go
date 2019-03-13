package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// Join func
func Join(g *game.Game, request, origin string) Data {
	phase := game.Joining
	find := func(p *player.Player) bool { return p.IsNameEmpty() }
	do := func(p *player.Player) error {
		data := strings.Split(request, "#")
		name := data[1]
		p.Join(name, origin)
		return nil
	}
	nextPlayerOperator := nextPlayer
	nextPhasePredicate := joinNextPhase
	playerPredicate := func(p *player.Player) bool { return p.IsNameEmpty() }
	return Data{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func joinNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.Count(searchCriteria) == 0
}

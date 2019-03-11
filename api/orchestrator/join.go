package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func (g *Game) join(request, origin string) (all []display.Info, me []display.Info, err error) {
	playerInTurn := g.players[g.playerInTurn]
	info := joinData(request, origin)
	return g.Info(), playerInTurn.Info(), g.playPhase(info)
}

func joinData(request, origin string) phaseData {
	phase := joining
	find := func(p *player.Player) bool { return p.IsNameEmpty() }
	do := func(p *player.Player) error { return joinAction(p, request, origin) }
	nextPlayerOperator := nextPlayer
	nextPhasePredicate := joinNextPhase
	playerPredicate := func(p *player.Player) bool { return p.IsNameEmpty() }
	return phaseData{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

func joinAction(p *player.Player, request, origin string) error {
	data := strings.Split(request, "#")
	name := data[1]
	p.Join(name, origin)
	return nil
}

func joinNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.Count(searchCriteria) == 0
}

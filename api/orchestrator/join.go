package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func (g *Game) join(request, origin string) (all []display.Info, me []display.Info, err error) {
	playerInTurn := g.playerInTurn
	info := g.joinData(request, origin)
	return g.Info(), g.players[playerInTurn].Info(), g.playPhase(info)
}

func (g *Game) joinData(request, origin string) dataPhase {
	phase := joining
	find := isNameEmpty
	do := func(p *player.Player) error { return joinAction(p, request, origin) }
	nextPlayerOperator := nextPlayer
	nextPhasePredicate := func() bool { return joinNextPhase(g.players, isNameEmpty) }
	return dataPhase{phase, find, do, nextPlayerOperator, nextPhasePredicate}
}

func joinAction(p *player.Player, request, origin string) error {
	data := strings.Split(request, "#")
	name := data[1]
	p.Join(name, origin)
	return nil
}

func nextPlayer(playerInTurn uint8) uint8 { return (playerInTurn + 1) % 5 }

func joinNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return players.Count(searchCriteria) == 0
}

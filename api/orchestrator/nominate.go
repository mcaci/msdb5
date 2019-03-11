package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func (g *Game) nominate(request, origin string) (all []display.Info, me []display.Info, err error) {
	playerInTurn := g.playerInTurn
	info := g.nominateData(request, origin)
	return g.Info(), g.players[playerInTurn].Info(), g.playPhase(info)
}

func (g *Game) nominateData(request, origin string) phaseData {
	phase := companionChoice
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) error {
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		return g.setCompanion(c)
	}
	nextPlayerOperator := func(playerInTurn uint8) uint8 { return playerInTurn }
	nextPhasePredicate := nominateNextPhase
	return phaseData{phase, find, do, nextPlayerOperator, nextPhasePredicate, nil}
}

func nominateNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return true
}

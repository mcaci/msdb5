package action

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/api/game"
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

// NominateData func
func NominateData(g *game.Game, request, origin string) Data {
	phase := game.CompanionChoice
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g.PlayerInTurn(), origin) }
	do := func(p *player.Player) error {
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		if err != nil {
			return err
		}
		pl, err := g.Players().Find(func(p *player.Player) bool { return p.Has(c) })
		if err != nil {
			return err
		}
		g.SetCompanion(c, pl)
		return nil
	}
	nextPlayerOperator := func(playerInTurn uint8) uint8 { return playerInTurn }
	nextPhasePredicate := nominateNextPhase
	return Data{phase, find, do, nextPlayerOperator, nextPhasePredicate, nil}
}

func nominateNextPhase(players playerset.Players, searchCriteria func(*player.Player) bool) bool {
	return true
}

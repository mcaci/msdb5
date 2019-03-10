package orchestrator

import (
	"errors"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) nominate(request, origin string) (all []display.Info, me []display.Info, err error) {
	data := strings.Split(request, "#")
	action := data[0]
	number := data[1]
	seed := data[2]
	playerInTurn := g.playerInTurn
	if action == "Companion" {
		find := func(p *player.Player) bool { return isActive(g, p, origin) }
		do := func(p *player.Player) (err error) {
			c, err := card.Create(number, seed)
			if err != nil {
				return
			}
			return g.setCompanion(c)
		}
		nextPlayerSupplier := func() uint8 { return g.playerInTurn }
		nextPhasePredicate := func() bool { return true }
		err = g.playPhase(companionChoice, find, do, nextPlayerSupplier, nextPhasePredicate)
		return g.Info(), g.players[playerInTurn].Info(), err
	}
	return g.Info(), g.players[playerInTurn].Info(), errors.New("COMPANION action not invoked")
}

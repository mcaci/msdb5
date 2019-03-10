package orchestrator

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Nominate func
func (g *Game) Nominate(number, seed, origin string) (err error) {
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
	return g.playPhase(companionChoice, find, do, nextPlayerSupplier, nextPhasePredicate)
}

func (g *Game) nominate(action, number, seed, origin string) (err error) {
	if action == "Companion" {
		return g.Nominate(number, seed, origin)
	}
	return errors.New("COMPANION action not invoked")
}

package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Nominate func
func (g *Game) Nominate(number, seed, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	nextPlayerSupplier := func() uint8 { return g.playerInTurn }
	nextPhasePredicate := func() bool { return true }
	return g.nominate(companionChoice, number, seed, origin, find, nextPlayerSupplier, nextPhasePredicate)
}

func (g *Game) nominate(phase phase, number, seed, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	_, err = g.players.Find(find)
	if err != nil {
		return
	}
	c, err := card.Create(number, seed)
	if err != nil {
		return
	}
	p, err := g.players.Find(func(p *player.Player) bool { return p.Has(c) })
	if err != nil {
		return
	}
	g.setCompanion(c, p)
	g.nextPlayer(nextPlayerSupplier)
	g.nextPhase(nextPhasePredicate)
	return
}

package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

// Join func
func (g *Game) Join(name, origin string) (err error) {
	find := isNameEmpty
	nextPlayerSupplier := func() uint8 { return 0 }
	nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
	return g.join(joining, name, origin, find, nextPlayerSupplier, nextPhasePredicate)
}

func (g *Game) join(phase phase, name, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
	if err != nil {
		return
	}
	p.Join(name, origin)
	g.nextPlayer(nextPlayerSupplier)
	g.nextPhase(nextPhasePredicate)
	return err
}

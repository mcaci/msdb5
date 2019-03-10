package orchestrator

import (
	"github.com/nikiforosFreespirit/msdb5/player"
)

// Join func
func (g *Game) Join(name, origin string) (err error) {
	find := isNameEmpty
	do := func(p *player.Player) error {
		p.Join(name, origin)
		return nil
	}
	nextPlayerSupplier := func() uint8 { return 0 }
	nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
	return g.playPhase(joining, find, do, nextPlayerSupplier, nextPhasePredicate)
}

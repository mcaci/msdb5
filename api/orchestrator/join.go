package orchestrator

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/player"
)

// Join func
func (g *Game) Join(name, origin string) (err error) {
	find := isNameEmpty
	do := func(p *player.Player) error {
		p.Join(name, origin)
		return nil
	}
	nextPlayerSupplier := func() uint8 { return (g.playerInTurn + 1) % 5 }
	nextPhasePredicate := func() bool { return g.players.Count(isNameEmpty) == 0 }
	return g.playPhase(joining, find, do, nextPlayerSupplier, nextPhasePredicate)
}

func (g *Game) join(action, name, origin string) (err error) {
	if action == "Join" {
		return g.Join(name, origin)
	}
	return errors.New("JOIN action not invoked")
}

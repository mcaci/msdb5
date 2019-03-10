package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

func (g *Game) playPhase(phase phase, find func(*player.Player) bool, do func(*player.Player) error, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
	if err != nil {
		return
	}
	err = do(p)
	if err != nil {
		return
	}
	g.nextPlayer(nextPlayerSupplier)
	g.nextPhase(nextPhasePredicate)
	return err
}

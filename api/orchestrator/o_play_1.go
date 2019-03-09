package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

func (g *Game) playRoundNotEnded(phase phase, number, seed, origin string, find func(*player.Player) bool, nextPlayerSupplier func() uint8, nextPhasePredicate func() bool) (err error) {
	if err = g.phaseCheck(phase); err != nil {
		return
	}
	p, err := g.players.Find(find)
	if err != nil {
		return
	}
	c, err := p.Play(number, seed)
	if err != nil {
		return
	}
	g.info.PlayedCards().Add(c)
	g.nextPlayer(nextPlayerSupplier)
	g.nextPhase(nextPhasePredicate)
	return
}

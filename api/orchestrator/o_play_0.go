package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

// Play func
func (g *Game) Play(number, seed, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	var nextPlayerSupplier (func() uint8)
	nextPhasePredicate := func() bool { return verifyEndGame(g) }
	roundMayEnd := len(*g.info.PlayedCards()) >= 4
	if roundMayEnd {
		return g.playRoundEnded(playBriscola, number, seed, origin, find, nextPlayerSupplier, nextPhasePredicate)
	}
	nextPlayerSupplier = func() uint8 { return (g.playerInTurn + 1) % 5 }
	return g.playRoundNotEnded(playBriscola, number, seed, origin, find, nextPlayerSupplier, nextPhasePredicate)
}

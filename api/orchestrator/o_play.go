package orchestrator

import "github.com/nikiforosFreespirit/msdb5/player"

// Play func
func (g *Game) Play(number, seed, origin string) (err error) {
	find := func(p *player.Player) bool { return isActive(g, p, origin) }
	var nextPlayerSupplier (func() uint8)
	nextPhasePredicate := func() bool { return verifyEndGame(g) }
	roundMayEnd := len(*g.info.PlayedCards()) >= 4
	if roundMayEnd {
		do := func(p *player.Player) (err error) {
			c, err := p.Play(number, seed)
			if err != nil {
				return
			}
			g.info.PlayedCards().Add(c)
			winnerIndex := winner(g)
			g.players[winnerIndex].Collect(g.info.PlayedCards())
			return
		}
		nextPlayerSupplier = func() uint8 { return winner(g) }
		return g.playPhase(playBriscola, find, do, nextPlayerSupplier, nextPhasePredicate)
	}
	do := func(p *player.Player) (err error) {
		c, err := p.Play(number, seed)
		if err != nil {
			return
		}
		g.info.PlayedCards().Add(c)
		return
	}
	nextPlayerSupplier = func() uint8 { return (g.playerInTurn + 1) % 5 }
	return g.playPhase(playBriscola, find, do, nextPlayerSupplier, nextPhasePredicate)
}

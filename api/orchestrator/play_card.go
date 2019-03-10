package orchestrator

import (
	"errors"

	"github.com/nikiforosFreespirit/msdb5/player"
)

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
		nextPlayerSupplier = func() uint8 {
			winnerIndex := winner(g)
			g.info.PlayedCards().Clear()
			return winnerIndex
		}
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

func (g *Game) play(action, number, seed, origin string) (err error) {
	if action == "Card" {
		return g.Play(number, seed, origin)
	}
	return errors.New("CARD action not invoked")
}

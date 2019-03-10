package orchestrator

import (
	"errors"
	"strings"

	"github.com/nikiforosFreespirit/msdb5/display"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) play(request, origin string) (all []display.Info, me []display.Info, err error) {
	data := strings.Split(request, "#")
	action := data[0]
	number := data[1]
	seed := data[2]
	playerInTurn := g.playerInTurn
	if action == "Card" {
		// err = g.Play(number, seed, origin)
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
			err = g.playPhase(playBriscola, find, do, nextPlayerSupplier, nextPhasePredicate)
		} else {
			do := func(p *player.Player) (err error) {
				c, err := p.Play(number, seed)
				if err != nil {
					return
				}
				g.info.PlayedCards().Add(c)
				return
			}
			nextPlayerSupplier = func() uint8 { return (g.playerInTurn + 1) % 5 }
			err = g.playPhase(playBriscola, find, do, nextPlayerSupplier, nextPhasePredicate)
		}
		return g.Info(), g.players[playerInTurn].Info(), err
	}
	return g.Info(), g.players[playerInTurn].Info(), errors.New("CARD action not invoked")
}

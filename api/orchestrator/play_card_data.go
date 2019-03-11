package orchestrator

import (
	"strings"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func (g *Game) playData(request, origin string) phaseData {
	phase := playBriscola
	find := func(p *player.Player) bool { return isExpectedPlayer(p, g, origin) }
	do := func(p *player.Player) (err error) {
		data := strings.Split(request, "#")
		number := data[1]
		seed := data[2]
		c, err := card.Create(number, seed)
		p.Play(c)
		g.info.PlayedCards().Add(c)
		return
	}
	nextPlayerOperator := nextPlayer
	nextPhasePredicate := g.endGameCondition
	playerPredicate := func(p *player.Player) bool { return p.IsHandEmpty() }
	return phaseData{phase, find, do, nextPlayerOperator, nextPhasePredicate, playerPredicate}
}

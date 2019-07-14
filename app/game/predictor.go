package game

import (
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type predictor interface {
	Caller() *player.Player
	Companion() *player.Player
	Players() team.Players
	Briscola() card.ID
}

func anyoneHasAllBriscola(g predictor, limit uint8) (bool, bool) {
	highbriscolaCard := briscola.Serie(g.Briscola())
	var callers, others bool
	var roundsChecked uint8
	for _, card := range highbriscolaCard {
		if roundsChecked == limit {
			break
		}
		_, p := g.Players().Find(func(p *player.Player) bool { return p.Has(card) })
		if p == nil { // no one has card
			continue
		}
		if p == g.Caller() || p == g.Companion() {
			callers = true
		} else {
			others = true
		}
		if callers == others {
			break
		}
		roundsChecked++
	}
	return callers, others
}

package game

import (
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type checker interface {
	Players() team.Players
}

func check(g checker) bool {
	return team.Count(g.Players(), player.IsHandEmpty) == 5
}

type predictor interface {
	Briscola() card.ID
	Caller() *player.Player
	Companion() *player.Player
	IsRoundOngoing() bool
	Players() team.Players
}

func predict(g predictor, roundsBefore, limit uint8) bool {
	return !g.IsRoundOngoing() && roundsBefore <= limit && oneTeamHasAllBriscola(g, limit)
}

func oneTeamHasAllBriscola(g predictor, limit uint8) bool {
	highbriscolaCard := briscola.Serie(g.Briscola())
	var callers, others bool
	var roundsChecked uint8
	for _, card := range highbriscolaCard {
		if roundsChecked == limit {
			break
		}
		_, p := g.Players().Find(player.IsCardInHand(card))
		if p == nil { // no one has card
			continue
		}
		isPlayerInCallers := p == g.Caller() || p == g.Companion()
		callers = callers || isPlayerInCallers
		others = others || !isPlayerInCallers
		if callers == others {
			break
		}
		roundsChecked++
	}
	return callers != others
}

package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/phase"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type phaseInformationProvider interface {
	Briscola() card.Item
	Caller() *player.Player
	Companion() *player.Player
	IsNewRoundToStart() bool
	IsSideUsed() bool
	Phase() phase.ID
	Players() team.Players
	ExchangeInput() string
}

// Phase func
func Phase(g phaseInformationProvider) phase.ID {
	current := g.Phase()
	shouldChange := phaseShouldChange(g)
	if shouldChange && current == phase.InsideAuction && !g.IsSideUsed() {
		current++
	}
	if shouldChange {
		current++
	}
	return current
}

func phaseShouldChange(g phaseInformationProvider) bool {
	isNext := true
	switch g.Phase() {
	case phase.Joining:
		isNext = team.Count(g.Players(), player.IsNameEmpty) == 0
	case phase.InsideAuction:
		isNext = team.Count(g.Players(), player.Folded) == 4
	case phase.ExchangingCards:
		isNext = g.ExchangeInput() == "0"
	case phase.PlayingCards:
		roundsBefore := uint8(len(*g.Players()[0].Hand()))
		const limit = 3
		isNext = predict(g, roundsBefore, limit) || checkAllWithEmptyHands(g)
	}
	return isNext
}

func checkAllWithEmptyHands(g interface{ Players() team.Players }) bool {
	return team.Count(g.Players(), player.IsHandEmpty) == 5
}

func predict(g phaseInformationProvider, roundsBefore, limit uint8) bool {
	highbriscolaCard := briscola.Serie(g.Briscola())
	var callersHave, othersHave bool
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
		callersHave = callersHave || isPlayerInCallers
		othersHave = othersHave || !isPlayerInCallers
		if callersHave == othersHave {
			break
		}
		roundsChecked++
	}
	return g.IsNewRoundToStart() && roundsBefore <= limit && callersHave != othersHave
}

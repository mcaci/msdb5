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
		const limit = 5
		roundsBefore := uint8(len(*g.Players()[0].Hand()))
		playedAllCards := team.Count(g.Players(), player.Folded) == 4
		isNext = (g.IsNewRoundToStart() && predict(g, roundsBefore, limit)) || playedAllCards
	}
	return isNext
}

func predict(g phaseInformationProvider, roundsBefore, limit uint8) bool {
	highbriscolaCard := briscola.Serie(g.Briscola())
	var callersHave, othersHave bool
	var cardsChecked uint8
	for _, card := range highbriscolaCard {
		if cardsChecked == limit {
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
		cardsChecked++
	}
	return cardsChecked == limit && callersHave != othersHave
}

package phase

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/briscola"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type phaseInformationProvider interface {
	Briscola() card.Item
	Caller() *player.Player
	Companion() *player.Player
	IsRoundOngoing() bool
	IsSideUsed() bool
	Players() team.Players
	ExchangeInput() string
}

// Next func
func (current ID) Next(g phaseInformationProvider) ID {
	if !current.phaseShouldChange(g) {
		return current
	}
	if current != InsideAuction || g.IsSideUsed() {
		return current + 1
	}
	return current + 2
}

func (current ID) phaseShouldChange(g phaseInformationProvider) bool {
	isNext := true
	switch current {
	case Joining:
		isNext = team.Count(g.Players(), player.IsNameEmpty) == 0
	case InsideAuction:
		isNext = team.Count(g.Players(), player.Folded) == 4
	case ExchangingCards:
		isNext = g.ExchangeInput() == "0"
	case PlayingCards:
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
	return !g.IsRoundOngoing() && roundsBefore <= limit && callersHave != othersHave
}

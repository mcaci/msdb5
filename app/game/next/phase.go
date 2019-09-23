package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/phase"
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
		isNext = g.Players().None(player.IsNameEmpty)
	case phase.InsideAuction:
		isNext = team.Count(g.Players(), player.Folded) == 4
	case phase.ExchangingCards:
		isNext = g.ExchangeInput() == "0"
	case phase.PlayingCards:
		isNext = isAnticipatedEnd(g) || g.Players().All(player.IsHandEmpty)
	}
	return isNext
}

func isAnticipatedEnd(g phaseInformationProvider) bool {
	var isAnticipatedEnd bool
	const limit = 3
	roundsBefore := uint8(len(*g.Players()[0].Hand()))
	if roundsBefore <= limit {
		isAnticipatedEnd = g.IsNewRoundToStart() && predict(g, roundsBefore)
	}
	return isAnticipatedEnd
}

func predict(g phaseInformationProvider, roundsBefore uint8) bool {
	highbriscolaCard := briscola.Serie(g.Briscola().Seed())
	var teams [2]bool
	var cardsChecked uint8
	for _, card := range highbriscolaCard {
		if g.Players().None(player.IsCardInHand(card)) { // no one has card
			continue
		}
		_, p := g.Players().Find(player.IsCardInHand(card))
		matching := player.Matching(p)
		isPlayerInTeam1 := matching(g.Caller()) || matching(g.Companion())
		teams[0] = teams[0] || isPlayerInTeam1
		teams[1] = teams[1] || !isPlayerInTeam1
		if teams[0] == teams[1] {
			return false
		}
		cardsChecked++
		if cardsChecked == roundsBefore {
			return true
		}
	}
	return false
}

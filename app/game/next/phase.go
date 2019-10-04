package next

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/phase"
	"github.com/mcaci/msdb5/dom/player"
	"github.com/mcaci/msdb5/dom/team"
)

type phaseInformationProvider interface {
	Briscola() card.Item
	IsNewRoundToStart() bool
	IsSideUsed() bool
	Phase() phase.ID
	Players() team.Players
	team.Callers
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
	highbriscolaCard := serie(g.Briscola().Seed())
	var teams [2]bool
	var cardsChecked uint8
	for _, card := range highbriscolaCard {
		i, err := g.Players().Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.Players().At(i)
		isPlayerInTeam1 := team.IsInCallers(g, p)
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

func serie(briscola card.Seed) set.Cards {
	serie := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	cards := make(set.Cards, len(serie))
	for i, id := range serie {
		cards[i] = *card.MustID(id + 10*uint8(briscola))
	}
	return cards
}

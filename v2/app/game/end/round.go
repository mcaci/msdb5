package end

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

type Opts struct {
	PlayedCards  set.Cards
	Players      team.Players
	BriscolaCard interface{ Seed() card.Seed }
	Callers      briscola5.Callerer
}

func Cond(g *Opts) bool {
	if g.Players.All(player.EmptyHanded) {
		return true
	}
	isNewRoundToStart := len(g.PlayedCards) == 5
	if !isNewRoundToStart {
		return false
	}
	const limit = 3
	roundsToEnd := len(*g.Players[0].Hand())
	if roundsToEnd > limit {
		return false
	}
	var teams [2]bool
	var cardsChecked int
	for _, card := range serie(g.BriscolaCard.Seed()) {
		i, err := g.Players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.Players[i]
		isPlayerInCallersTeam := briscola5.IsInCallers(g.Callers)(p)
		teams[0] = teams[0] || isPlayerInCallersTeam
		teams[1] = teams[1] || !isPlayerInCallersTeam
		if teams[0] == teams[1] {
			return false
		}
		cardsChecked++
		if cardsChecked == roundsToEnd {
			return true
		}
	}
	return false
}

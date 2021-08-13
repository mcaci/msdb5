package end

import (
	"github.com/mcaci/msdb5/v2/app/briscola5"
	"github.com/mcaci/msdb5/v2/app/player"
	"github.com/mcaci/msdb5/v2/dom/briscola"
)

type Opts struct {
	PlayedCards  briscola.PlayedCards
	Players      player.Players
	Callers      briscola5.Callerer
	BriscolaCard briscola.Card
}

func Cond(g *Opts) bool {
	// no more cards to play
	if g.Players.All(player.EmptyHanded) {
		return true
	}
	isNewRoundToStart := len(*g.PlayedCards.Cards) == 5
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
	isPlayerInCallersTeamF := player.IsInCallers(g.Callers)
	for _, card := range briscola.Serie(g.BriscolaCard) {
		i, err := g.Players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.Players[i]
		isPlayerInCallersTeam := isPlayerInCallersTeamF(p)
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

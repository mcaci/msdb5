package end

import (
	"github.com/mcaci/msdb5/v2/dom/briscola"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
)

type Opts struct {
	PlayedCards  briscola.PlayedCards
	Players      briscola5.Players
	BriscolaCard briscola.Card
}

func Cond(g *Opts) bool {
	// no more cards to play
	if briscola5.ToGeneralPlayers(g.Players).All(player.EmptyHanded) {
		return true
	}
	isNewRoundToStart := len(*g.PlayedCards.Cards) == 5
	if !isNewRoundToStart {
		return false
	}
	const limit = 3
	roundsToEnd := len(*g.Players.At(0).Hand())
	if roundsToEnd > limit {
		return false
	}
	var teams [2]bool
	var cardsChecked int
	for _, card := range briscola.Serie(g.BriscolaCard) {
		i, err := briscola5.ToGeneralPlayers(g.Players).Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.Players.At(int(i))
		isPlayerInCallersTeam := briscola5.IsInCallers(&g.Players)(&p.Player)
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

package end

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Cond(g struct {
	PlayedCards  set.Cards
	Players      team.Players
	BriscolaCard interface{ Seed() card.Seed }
	Callers      briscola5.Callerer
}) bool {
	return g.Players.All(player.EmptyHanded) ||
		isAnticipatedEnd_v2(struct {
			players      team.Players
			playedCards  set.Cards
			briscolaCard interface{ Seed() card.Seed }
			callers      briscola5.Callerer
		}{players: g.Players, playedCards: g.PlayedCards, briscolaCard: g.BriscolaCard, callers: g.Callers})
}

func isAnticipatedEnd_v2(g struct {
	players      team.Players
	playedCards  set.Cards
	briscolaCard interface{ Seed() card.Seed }
	callers      briscola5.Callerer
}) bool {
	var isAnticipatedEnd bool
	const limit = 3
	roundsBefore := uint8(len(*g.players[0].Hand()))
	if roundsBefore <= limit {
		isNewRoundToStart := len(g.playedCards) == 5
		isAnticipatedEnd = isNewRoundToStart && predict_v2(struct {
			players      team.Players
			briscolaCard interface{ Seed() card.Seed }
			callers      briscola5.Callerer
		}{
			players: g.players, briscolaCard: g.briscolaCard, callers: g.callers,
		}, roundsBefore)
	}
	return isAnticipatedEnd
}

func predict_v2(g struct {
	players      team.Players
	briscolaCard interface{ Seed() card.Seed }
	callers      briscola5.Callerer
}, roundsBefore uint8) bool {
	highbriscolaCard := serie(g.briscolaCard.Seed())
	var teams [2]bool
	var cardsChecked uint8
	for _, card := range highbriscolaCard {
		i, err := g.players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.players[i]
		isPlayerInCallersTeam := briscola5.IsInCallers(g.callers)(p)
		teams[0] = teams[0] || isPlayerInCallersTeam
		teams[1] = teams[1] || !isPlayerInCallersTeam
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

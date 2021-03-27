package end

import (
	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Cond(g struct {
	playedCards       set.Cards
	players           team.Players
	briscolaCard      card.Item
	caller, companion *player.Player
}) bool {
	// next phase
	return g.players.All(player.EmptyHanded) ||
		isAnticipatedEnd_v2(struct {
			players           team.Players
			playedCards       set.Cards
			briscolaCard      card.Item
			caller, companion *player.Player
		}{players: g.players, playedCards: g.playedCards, briscolaCard: g.briscolaCard,
			caller: g.caller, companion: g.companion})
}

func isAnticipatedEnd_v2(g struct {
	players           team.Players
	playedCards       set.Cards
	briscolaCard      card.Item
	caller, companion *player.Player
}) bool {
	var isAnticipatedEnd bool
	const limit = 3
	roundsBefore := uint8(len(*g.players[0].Hand()))
	if roundsBefore <= limit {
		isNewRoundToStart := len(g.playedCards) == 5
		isAnticipatedEnd = isNewRoundToStart && predict_v2(struct {
			players      team.Players
			briscolaCard card.Item
			caller       *player.Player
			companion    *player.Player
		}{
			players: g.players, briscolaCard: g.briscolaCard, caller: g.caller, companion: g.companion,
		}, roundsBefore)
	}
	return isAnticipatedEnd
}

func predict_v2(g struct {
	players           team.Players
	briscolaCard      card.Item
	caller, companion *player.Player
}, roundsBefore uint8) bool {
	highbriscolaCard := serie(g.briscolaCard.Seed())
	var teams [2]bool
	var cardsChecked uint8
	for _, card := range highbriscolaCard {
		i, err := g.players.Index(player.IsCardInHand(card))
		if err != nil { // no one has card
			continue
		}
		p := g.players.At(i)
		isPlayerInCallersTeam := team.IsInCallers(callers{caller: g.caller, companion: g.companion})(p)
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

func serie(briscola card.Seed) set.Cards {
	serie := []uint8{1, 3, 10, 9, 8, 7, 6, 5, 4, 2}
	cards := make(set.Cards, len(serie))
	for i, id := range serie {
		cards[i] = *card.MustID(id + 10*uint8(briscola))
	}
	return cards
}

type callers struct {
	caller, companion *player.Player
}

func (c callers) Caller() *player.Player    { return c.caller }
func (c callers) Companion() *player.Player { return c.companion }

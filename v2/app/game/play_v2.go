package game

import (
	"container/list"
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runPlay_v2(g struct {
	players           team.Players
	briscolaCard      card.Item
	lastPlaying       list.List
	caller, companion *player.Player
}) struct {
	onBoard set.Cards
} {
	var playedCards set.Cards

	for !endCond(struct {
		playedCards  set.Cards
		players      team.Players
		briscolaCard card.Item
		caller       *player.Player
		companion    *player.Player
	}{
		playedCards:  playedCards,
		players:      g.players,
		briscolaCard: g.briscolaCard,
		caller:       g.caller,
		companion:    g.companion,
	}) {
		pl := CurrentPlayer(g.lastPlaying)
		hnd := pl.Hand()
		if len(*hnd) > 0 {
			rand.Seed(time.Now().Unix())
			idx := rand.Intn(len(*hnd))
			crd := (*hnd)[idx]
			index := hnd.Find(crd)
			playedCards.Add((*hnd)[index])
			*hnd = append((*hnd)[:index], (*hnd)[index+1:]...)
		}

		// next player
		idx, err := CurrentPlayerIndex(pl, g.players)
		if err != nil {
			return struct{ onBoard set.Cards }{}
		}
		nextPlayer := roundRobin(idx, 1, numberOfPlayers)
		if !IsRoundOngoing(playedCards) {
			// end current round
			winningCardIndex := indexOfWinningCard(playedCards, g.briscolaCard.Seed())
			nextPlayer = roundRobin(nextPlayer, winningCardIndex, numberOfPlayers)

			set.Move(collect.NewRoundCards(&playedCards).Set(), g.players[nextPlayer].Pile())
		}
		track.Player(&g.lastPlaying, g.players[nextPlayer])
	}
	return struct{ onBoard set.Cards }{
		onBoard: playedCards,
	}
}

func endCond(g struct {
	playedCards       set.Cards
	players           team.Players
	briscolaCard      card.Item
	caller, companion *player.Player
}) bool {
	// next phase
	return g.players.All(player.IsHandEmpty) ||
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

type callers struct {
	caller, companion *player.Player
}

func (c callers) Caller() *player.Player    { return c.caller }
func (c callers) Companion() *player.Player { return c.companion }

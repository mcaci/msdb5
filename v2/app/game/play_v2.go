package game

import (
	"container/list"
	"errors"
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/app/game/end"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runPlay_v2(g struct {
	players      team.Players
	briscolaCard interface{ Seed() card.Seed }
	callers      team.Callers
}) struct {
	onBoard set.Cards
} {
	var lastPlaying list.List
	track.Player(&lastPlaying, g.callers.Caller())
	var playedCards set.Cards

	for !end.Cond(struct {
		PlayedCards  set.Cards
		Players      team.Players
		BriscolaCard interface{ Seed() card.Seed }
		Callers      team.Callers
	}{
		PlayedCards:  playedCards,
		Players:      g.players,
		BriscolaCard: g.briscolaCard,
		Callers:      g.callers,
	}) {
		pl := currentPlayer(lastPlaying)
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
		idx, err := currentPlayerIndex(pl, g.players)
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
		track.Player(&lastPlaying, g.players[nextPlayer])
	}
	return struct{ onBoard set.Cards }{
		onBoard: playedCards,
	}
}

func currentPlayer(l list.List) *player.Player { return l.Front().Value.(*player.Player) }
func currentPlayerIndex(cp *player.Player, pls team.Players) (uint8, error) {
	for i := range pls {
		if pls[i] != cp {
			continue
		}
		return uint8(i), nil
	}
	return 0, errors.New("Not found")
}

package play

import (
	"container/list"
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Round(g struct {
	Players      team.Players
	LastPlaying  list.List
	BriscolaCard interface{ Seed() card.Seed }
	PlayedCards  set.Cards
}) struct{ OnBoard set.Cards } {
	pl := currentPlayer(g.LastPlaying)
	hnd := pl.Hand()
	if len(*hnd) > 0 {
		rand.Seed(time.Now().Unix())
		idx := rand.Intn(len(*hnd))
		crd := (*hnd)[idx]
		index := hnd.Find(crd)
		g.PlayedCards.Add((*hnd)[index])
		*hnd = append((*hnd)[:index], (*hnd)[index+1:]...)
	}

	// next player
	idx, err := currentPlayerIndex(pl, g.Players)
	if err != nil {
		return struct{ OnBoard set.Cards }{}
	}
	nextPlayer := roundRobin(idx, 1, numberOfPlayers)
	if !isRoundOngoing(g.PlayedCards) {
		// end current round
		winningCardIndex := indexOfWinningCard(g.PlayedCards, g.BriscolaCard.Seed())
		nextPlayer = roundRobin(nextPlayer, winningCardIndex, numberOfPlayers)

		set.Move(collect.NewRoundCards(&g.PlayedCards).Set(), g.Players[nextPlayer].Pile())
	}
	track.Player(&g.LastPlaying, g.Players[nextPlayer])
	return struct{ OnBoard set.Cards }{
		OnBoard: g.PlayedCards,
	}
}

package play

import (
	"container/list"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/collect"
	"github.com/mcaci/msdb5/v2/app/game/end"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(g struct {
	Players      team.Players
	BriscolaCard interface{ Seed() card.Seed }
	Callers      briscola5.Callerer
}) struct {
	OnBoard set.Cards
} {
	var playedCards set.Cards
	plIdx, err := currentPlayerIndex(g.Callers.Caller(), g.Players)
	if err != nil {
		log.Fatal("didn't expect to arrive at this point")
	}

	for !end.Cond(&end.Opts{
		PlayedCards:  playedCards,
		Players:      g.Players,
		BriscolaCard: g.BriscolaCard,
		Callers:      g.Callers,
	}) {
		rand.Seed(time.Now().Unix())
		hnd := g.Players[plIdx].Hand()
		info := Round(&RoundOptions{
			PlHand:       *hnd,
			Idx:          uint8(rand.Intn(len(*hnd))),
			PlayedCards:  playedCards,
			NPlayers:     uint8(len(g.Players)),
			BriscolaCard: g.BriscolaCard,
		})
		playedCards = info.OnBoard
		plIdx = info.NextPl
		if !info.NextRnd {
			continue
		}
		set.Move(collect.NewRoundCards(&playedCards).Set(), g.Players[plIdx].Pile())
	}
	return struct{ OnBoard set.Cards }{
		OnBoard: playedCards,
	}
}

func isRoundOngoing(playedCards set.Cards) bool { return len(playedCards) < 5 }
func currentPlayer(l list.List) *player.Player  { return l.Front().Value.(*player.Player) }
func currentPlayerIndex(cp *player.Player, pls team.Players) (uint8, error) {
	for i := range pls {
		if pls[i] != cp {
			continue
		}
		return uint8(i), nil
	}
	return 0, errors.New("Not found")
}

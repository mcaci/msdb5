package play

import (
	"container/list"
	"errors"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/app/game/end"
	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(g struct {
	Players      team.Players
	BriscolaCard interface{ Seed() card.Seed }
	Callers      team.Callers
}) struct {
	OnBoard set.Cards
} {
	var lastPlaying list.List
	track.Player(&lastPlaying, g.Callers.Caller())
	var playedCards set.Cards

	for !end.Cond(struct {
		PlayedCards  set.Cards
		Players      team.Players
		BriscolaCard interface{ Seed() card.Seed }
		Callers      team.Callers
	}{
		PlayedCards:  playedCards,
		Players:      g.Players,
		BriscolaCard: g.BriscolaCard,
		Callers:      g.Callers,
	}) {
		playedCards = Round(struct {
			Players      team.Players
			LastPlaying  list.List
			BriscolaCard interface{ Seed() card.Seed }
			PlayedCards  set.Cards
		}{
			Players:      g.Players,
			LastPlaying:  lastPlaying,
			BriscolaCard: g.BriscolaCard,
			PlayedCards:  playedCards,
		}).OnBoard
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

const numberOfPlayers = 5

func roundRobin(idx, off, size uint8) uint8 {
	return (idx + off) % size
}

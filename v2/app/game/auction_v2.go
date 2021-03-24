package game

import (
	"container/list"
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/mcaci/msdb5/v2/app/track"
	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func runAuction_v2(g struct {
	players     team.Players
	lastPlaying list.List
}, listenFor func(context.Context, func())) struct {
	score  auction.Score
	caller *player.Player
} {
	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- 60 + rand.Intn(60) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	var caller *player.Player
	var curr auction.Score

	for score := range numbers {
		next := auction.Score(score)
		res := checkUpdateScore_v2(struct {
			curr, next  auction.Score
			players     team.Players
			lastPlaying list.List
		}{
			curr:        curr,
			next:        next,
			players:     g.players,
			lastPlaying: g.lastPlaying,
		})
		curr = res.score
		if res.fold != nil {
			res.fold()
		}
		if team.Count(g.players, notFolded) == 1 {
			// set caller
			caller = g.players.At(g.players.MustIndex(notFolded))
			// next phase
			done <- struct{}{}
			close(done)
		}
		pl := CurrentPlayer(g.lastPlaying)
		idx, err := CurrentPlayerIndex(pl, g.players)
		if err != nil {
			log.Fatalf("error found: %v. Exiting.", err)
		}
		// next player
		nextPlayerInfo := rotate(idx, g.players)
		if nextPlayerInfo.err != nil {
			log.Fatalf("error found: %v. Exiting.", err)
		}
		track.Player(&g.lastPlaying, nextPlayerInfo.p)
	}
	return struct {
		score  auction.Score
		caller *player.Player
	}{
		score:  curr,
		caller: caller,
	}
}

func rotate(currID uint8, players team.Players) struct {
	p   *player.Player
	err error
} {
	for i := 0; i < 2*len(players); i++ {
		currID = roundRobin(currID, 1, uint8(len(players)))
		if player.Folded(players[currID]) {
			continue
		}
		return struct {
			p   *player.Player
			err error
		}{p: players[currID], err: nil}
	}
	return struct {
		p   *player.Player
		err error
	}{p: nil, err: errors.New("rotated twice on the number of players and no player found in play.")}
}

func checkUpdateScore_v2(g struct {
	curr, next  auction.Score
	players     team.Players
	lastPlaying list.List
}) struct {
	score auction.Score
	fold  func()
} {
	toFold := player.Folded(CurrentPlayer(g.lastPlaying)) || !auction.ScoreCmp(g.curr, g.next)
	if toFold {
		return struct {
			score auction.Score
			fold  func()
		}{
			score: g.curr,
			fold:  func() { CurrentPlayer(g.lastPlaying).Fold() },
		}
	}
	newScore := auction.Max120(g.curr, g.next)
	return struct {
		score auction.Score
		fold  func()
	}{
		score: newScore,
		fold: func() {
			if newScore >= 120 {
				for _, p := range g.players {
					if p == CurrentPlayer(g.lastPlaying) {
						continue
					}
					p.Fold()
				}
			}
		},
	}
}

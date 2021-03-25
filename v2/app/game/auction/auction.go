package auction

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

type Options struct {
	Players     team.Players
	LastPlaying list.List
}

func Run(opt *Options, listenFor func(context.Context, func())) struct {
	Score  auction.Score
	Caller *player.Player
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
		pl := currentPlayer(opt.LastPlaying)
		next := auction.Score(score)

		// Player is folding
		if toFold := player.Folded(pl) || !auction.ScoreCmp(curr, next); toFold {
			pl.Fold()
		}

		// Fold everyone if score is 120 or more
		curr = auction.Max120(curr, next)
		if curr >= 120 {
			for _, p := range opt.Players {
				if p == pl {
					continue
				}
				p.Fold()
			}
		}

		// End the loop if only one person is left
		if team.Count(opt.Players, notFolded) == 1 {
			// set caller
			caller = opt.Players.At(opt.Players.MustIndex(notFolded))
			// next phase
			done <- struct{}{}
			close(done)
		}

		// next player
		currID, err := currentPlayerIndex(opt.LastPlaying, opt.Players)
		if err != nil {
			log.Fatalf("error found: %v. Exiting.", err)
		}
		var found bool
		for i := 0; i < 2*len(opt.Players); i++ {
			currID = roundRobin(currID, 1, uint8(len(opt.Players)))
			if player.Folded(opt.Players[currID]) {
				continue
			}
			found = true
			track.Player(&opt.LastPlaying, opt.Players[currID])
		}
		if !found {
			log.Fatalln("rotated twice on the number of players and no player found in play. Exiting.")
		}
	}
	return struct {
		Score  auction.Score
		Caller *player.Player
	}{
		Score:  curr,
		Caller: caller,
	}
}

func notFolded(p *player.Player) bool          { return !player.Folded(p) }
func roundRobin(idx, off, size uint8) uint8    { return (idx + off) % size }
func currentPlayer(l list.List) *player.Player { return l.Front().Value.(*player.Player) }
func currentPlayerIndex(l list.List, pls team.Players) (uint8, error) {
	cp := currentPlayer(l)
	for i := range pls {
		if pls[i] != cp {
			continue
		}
		return uint8(i), nil
	}
	return 0, errors.New("Not found")
}

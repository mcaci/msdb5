package auction

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func Run(players briscola5.Players, listenFor func(context.Context, func())) struct {
	Score  auction.Score
	Caller *briscola5.Player
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

	var score auction.Score
	var currID uint8

	for n := range numbers {
		r := Round(struct {
			curr, prop auction.Score
			currID     uint8
			players    briscola5.Players
		}{
			curr:    score,
			prop:    auction.Score(n),
			currID:  currID,
			players: players,
		})
		score = r.s
		currID = r.id
		if !r.end {
			continue
		}
		done <- struct{}{}
		close(done)
	}
	return struct {
		Score  auction.Score
		Caller *briscola5.Player
	}{
		Score:  score,
		Caller: players[players.MustIndex(notFolded)],
	}
}

func notFolded(p *briscola5.Player) bool { return !briscola5.Folded(p) }
func mustRotateOnNotFolded(players briscola5.Players, from uint8) uint8 {
	id, err := rotateOn(players, from, notFolded)
	if err != nil {
		log.Fatalf("error found: %v. Exiting.", err)
	}
	return id
}
func rotateOn(players briscola5.Players, idx uint8, appliesTo briscola5.Predicate) (uint8, error) {
	for i := 0; i < 2*len(players); i++ {
		idx = (idx + 1) % uint8(len(players))
		if !appliesTo(players[idx]) {
			continue
		}
		return idx, nil
	}
	return 0, fmt.Errorf("rotated twice on the number of players and no player found in play.")
}

package auction

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/mcaci/msdb5/v2/dom/auction"
	"github.com/mcaci/msdb5/v2/dom/player"
	"github.com/mcaci/msdb5/v2/dom/team"
)

func Run(players team.Players, listenFor func(context.Context, func())) struct {
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

	var score auction.Score
	var currID uint8

	for n := range numbers {
		r := Round(score, auction.Score(n), currID, players)
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
		Caller *player.Player
	}{
		Score:  score,
		Caller: players[players.MustIndex(notFolded)],
	}
}

func notFolded(p *player.Player) bool { return !player.Folded(p) }
func mustRotateOnNotFolded(players team.Players, from uint8) uint8 {
	id, err := rotateOn(players, from, notFolded)
	if err != nil {
		log.Fatalf("error found: %v. Exiting.", err)
	}
	return id
}
func rotateOn(players team.Players, idx uint8, appliesTo player.Predicate) (uint8, error) {
	for i := 0; i < 2*len(players); i++ {
		idx = (idx + 1) % uint8(len(players))
		if !appliesTo(players[idx]) {
			continue
		}
		return idx, nil
	}
	return 0, fmt.Errorf("rotated twice on the number of players and no player found in play.")
}

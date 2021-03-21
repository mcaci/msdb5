package game

import (
	"container/list"
	"context"
	"math/rand"

	"github.com/mcaci/ita-cards/set"
)

func runExchange_v2(g struct {
	opts        *Options
	side        set.Cards
	lastPlaying list.List
}, listenFor func(context.Context, func())) {
	if !g.opts.WithSide {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- rand.Intn(len(*CurrentPlayer(g.lastPlaying).Hand())) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	for idx := range numbers {
		if idx > 2 {
			done <- struct{}{}
			close(done)
			break
		}
		hnd := CurrentPlayer(g.lastPlaying).Hand()
		discardedCard := (*hnd)[idx]
		sideCards := g.side
		(*hnd)[idx] = sideCards[0]
		sideCards = append(sideCards[1:], discardedCard)
	}
}

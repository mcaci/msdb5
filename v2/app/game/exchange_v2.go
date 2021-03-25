package game

import (
	"context"
	"math/rand"

	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/v2/dom/player"
)

func runExchange_v2(g struct {
	opts *Options
	side set.Cards
	pl   *player.Player
}, listenFor func(context.Context, func())) {
	if !g.opts.WithSide {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- rand.Intn(len(*g.pl.Hand())) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	for idx := range numbers {
		if idx > 2 {
			done <- struct{}{}
			close(done)
		}
		hnd := g.pl.Hand()
		discardedCard := (*hnd)[idx]
		sideCards := g.side
		(*hnd)[idx] = sideCards[0]
		sideCards = append(sideCards[1:], discardedCard)
	}
}

package game

import (
	"context"
	"math/rand"
)

func runExchange(g *Game, listenFor func(context.Context, func())) {
	if !g.opts.WithSide {
		g.phase++
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- rand.Intn(len(*g.CurrentPlayer().Hand())) })
	go func() {
		<-done
		cancel()
		close(numbers)
	}()

	for idx := range numbers {
		if idx > 2 {
			g.phase++
			done <- struct{}{}
			close(done)
		}
		hnd := g.CurrentPlayer().Hand()
		discardedCard := (*hnd)[idx]
		sideCards := g.side
		(*hnd)[idx] = sideCards[0]
		sideCards = append(sideCards[1:], discardedCard)
	}
}

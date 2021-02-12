package game

import (
	"context"
	"math/rand"
)

func runExchange(g *Game, listenFor func(context.Context, chan<- int, func() int)) {
	if !g.withSide {
		g.phase++
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, numbers, func() int { return rand.Intn(len(*g.CurrentPlayer().Hand())) })
	go func() {
		<-done
		cancel()
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

package exchange

import (
	"context"
	"math/rand"

	"github.com/mcaci/ita-cards/set"
)

func Run(g struct {
	Hand, Side *set.Cards
}, listenFor func(context.Context, func())) {

	ctx, cancel := context.WithCancel(context.Background())
	numbers := make(chan int)
	done := make(chan struct{})
	go listenFor(ctx, func() { numbers <- rand.Intn(len(*g.Hand)) })
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
		Round(struct {
			Hand, Side *set.Cards
			hIdx, sIdx int
		}{
			Hand: g.Hand, Side: g.Side,
			hIdx: idx, sIdx: 0,
		})
	}
}

package frw

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	iterations = 5
	nMillis    = 1
)

func init() {
	rand.Seed(time.Now().Unix())
}

func WithAINames(names chan<- string) {
	ticker := time.NewTicker(nMillis * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < iterations; i++ {
		<-ticker.C
		names <- fmt.Sprintf("Player%d", i+1)
	}
	close(names)
}

func WithRand(ctx context.Context, numbers chan<- int, n func() int) {
	ticker := time.NewTicker(nMillis * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			numbers <- n()
		case <-ctx.Done():
			close(numbers)
			return
		}
	}
}

type TickerRand struct {
	context.Context
	tick *time.Ticker
	fCh  chan func()
}

func (tr *TickerRand) Input() <-chan func() {
	return tr.fCh
}

func (tr *TickerRand) StartExample(f func()) {
	ctx, cancel := context.WithCancel(tr.Context)
	defer cancel()
	for i := 0; i < 5; i++ {
		select {
		case <-tr.tick.C:
			tr.fCh <- f
		case <-ctx.Done():
			close(tr.fCh)
		}
	}
}

func WithRand2(a interface {
	context.Context
	Input() <-chan func()
}) {
	for {
		select {
		case f := <-a.Input():
			f()
		case <-a.Done():
			return
		}
	}
}

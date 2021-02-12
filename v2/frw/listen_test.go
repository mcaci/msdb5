package frw

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestWithAINames(t *testing.T) { // add -race to go test for running this test
	ch := make(chan string)
	go WithAINames(ch)
	for c := range ch {
		t.Log(c)
	}
}

func TestWithRand(t *testing.T) { // add -race to go test for running this test
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go WithRand(ctx, ch, func() int { return 0 })
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	for c := range ch {
		t.Log(c)
	}
}

func TestWithRand2(t *testing.T) { // add -race to go test for running this test
	tr := &TickerRand{Context: context.Background(), tick: time.NewTicker(10 * time.Millisecond), fCh: make(chan func())}
	go WithRand2(tr)
	tr.StartExample(func() {
		log.Println("hello")
	})
}

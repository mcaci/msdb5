package srv

import (
	"context"
	"log"
	"time"
)

type WaitSession struct {
	ready chan struct{}
}

var waitSession = WaitSession{ready: make(chan struct{})}

func Wait() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tick := time.NewTicker(1 * time.Second)
	var i int
waiter:
	for {
		select {
		case <-tick.C:
			i++
			log.Println(i)
		case <-waitSession.ready:
			break waiter
		case <-ctx.Done():
			break waiter
		}
	}
}

func SigAllJoin() {
	waitSession.ready <- struct{}{}
}

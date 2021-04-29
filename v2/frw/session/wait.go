package session

import (
	"context"
	"log"
	"time"
)

func Wait(on <-chan interface{}) {
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
		case <-on:
			break waiter
		case <-ctx.Done():
			break waiter
		}
	}
}

func Signal(on chan<- interface{}) {
	on <- struct{}{}
}

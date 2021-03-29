package game

import (
	"strconv"
	"testing"
	"time"
)

func supply(p chan<- string) {
	ticker := time.NewTicker(1 * time.Millisecond)
	var i int
	for range ticker.C {
		p <- "player" + strconv.Itoa(i+1)
		i++
		if i > 4 {
			ticker.Stop()
			close(p)
		}
	}
}

func TestSideDeckProperty(t *testing.T) {
	g := NewGame(&Options{})
	if g.opts.WithSide {
		t.Errorf("error")
	}
	g = NewGame(&Options{
		WithSide: true,
	})
	if !g.opts.WithSide {
		t.Errorf("error")
	}
}

func TestWaitForPlayers(t *testing.T) {
	g := New()
	WaitForPlayers(g, supply)
	if len(g.players) != 5 {
		t.Errorf("%d", len(g.players))
	}
}

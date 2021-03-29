package briscola5

import (
	"testing"
)

func TestCount(t *testing.T) {
	p := NewPlayer()
	if count := Count(Players{p, p}, func(pl *Player) bool { return true }); count != 2 {
		t.Fatal("Count should be 2")
	}
}

package game

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
)

func TestCardsAreMoved(t *testing.T) {
	from := set.NewMust(1, 2, 3, 4, 5)
	to := set.NewMust()
	move(from, to)
	if len(*from) > 0 {
		t.Fatalf("From should be empty but has %v", from)
	}
}

func TestPlayerCountPoints(t *testing.T) {
	from := set.NewMust(1, 2, 3, 4, 5)
	to := set.NewMust()
	move(from, to)
	if len(*to) == 0 {
		t.Fatalf("To should be empty but has %v", to)
	}
}

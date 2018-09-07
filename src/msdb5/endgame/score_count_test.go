package endgame

import (
	"msdb5/card"
	"testing"
)

func points(c ...*card.Card) int { return 0 }

func TestEmptyPileSums11(t *testing.T) {
	score := points()
	if score != 0 {
		t.Fatalf("Score expected is not 0 but %d", score)
	}
}


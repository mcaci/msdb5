package endgame

import (
	"msdb5/card"
	"testing"
)

func points(c []*card.Card) int { return 1 }

func TestEmptyPileSumsZero(t *testing.T) {
	score := points([]*card.Card{})
	if score != 0 {
		t.Fatalf("Score expected is not 0 but %d", score)
	}
}

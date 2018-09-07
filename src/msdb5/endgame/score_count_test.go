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

func TestPileWitOnehAceOnlySums11(t *testing.T) {
	ace, _ := card.ByID(1)
	cards := []*card.Card{ace}
	score := points(cards...)
	if score != 11 {
		t.Fatalf("Score expected is not 11 but %d", score)
	}
}

package endgame

import (
	"msdb5/card"
	"testing"
)

func TestEmptyPileSums0(t *testing.T) {
	score := CountPoints()
	if score != 0 {
		t.Fatalf("Score expected is not 0 but %d", score)
	}
}

func TestPileWitOnehAceOnlySums11(t *testing.T) {
	ace, _ := card.ByID(1)
	cards := []*card.Card{ace}
	score := CountPoints(cards...)
	if score != 11 {
		t.Fatalf("Score expected is not 11 but %d", score)
	}
}

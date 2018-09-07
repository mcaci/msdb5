package endgame

import (
	"msdb5/card"
	"testing"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards ...*card.Card) {
	score := CountPoints(cards...)
	if expectedScore != score {
		t.Fatalf("Score expected is not %d but %d", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWitOnehAceOnlySums11(t *testing.T) {
	ace, _ := card.ByID(1)
	testScoreCount(t, 11, ace)
}

package endgame

import (
	"msdb5/card"
	"testing"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards ...*card.Card) {
	score := CountPoints(cards...)
	if expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	ace, _ := card.ByID(1)
	testScoreCount(t, 11, ace)
}

func TestPileWithOneTwoOnehAceOnlySums11(t *testing.T) {
	two, _ := card.ByID(2)
	ace, _ := card.ByID(1)
	testScoreCount(t, 11, two, ace)
}

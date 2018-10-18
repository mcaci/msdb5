package score

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards ...uint8) {
	score := Compute(cards...)
	if expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, card.FillWithIDs(1)...)
}

func TestPileWithOneTwoOneAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, card.FillWithIDs(2, 1)...)
}

func TestPileWithOneAceOneTwoOneThreeSums21(t *testing.T) {
	testScoreCount(t, 21, card.FillWithIDs(1, 2, 3)...)
}
func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, card.FillWithIDs(card.Deck()...)...)
}

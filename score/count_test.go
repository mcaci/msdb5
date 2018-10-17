package score

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards ...card.Card) {
	score := Compute(cards...)
	if expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, card.Fill(card.WithIDs, 1)...)
}

func TestPileWithOneTwoOneAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, card.Fill(card.WithIDs, 2, 1)...)
}

func TestPileWithOneAceOneTwoOneThreeSums21(t *testing.T) {
	testScoreCount(t, 21, card.Fill(card.WithIDs, 1, 2, 3)...)
}
func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, card.Fill(card.WithIDs, deck.New().GetIDs()...)...)
}

package rule

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards card.Cards) {
	if score := Count(cards); expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0, nil)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, card.Cards{1})
}

func TestPileWithOneTwoOneAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, card.Cards{2, 1})
}

func TestPileWithOneAceOneTwoOneThreeSums21(t *testing.T) {
	testScoreCount(t, 21, card.Cards{1, 2, 3})
}
func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, card.Deck())
}

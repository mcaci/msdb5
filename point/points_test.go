package point

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/briscola"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards deck.Cards) {
	if score := Count(cards, briscola.Points); expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0, nil)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, deck.Cards{1})
}

func TestPileWithOneTwoOneAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, deck.Cards{2, 1})
}

func TestPileWithOneAceOneTwoOneThreeSums21(t *testing.T) {
	testScoreCount(t, 21, deck.Cards{1, 2, 3})
}
func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, deck.Deck())
}

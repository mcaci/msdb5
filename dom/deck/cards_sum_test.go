package deck

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards Cards) {
	if score := cards.Sum(func(card.ID) uint8 { return 1 }); expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0, nil)
}

func TestPileWithOneCardSums1(t *testing.T) {
	testScoreCount(t, 1, Cards{1})
}

func TestPileWithThreeCardsSums3(t *testing.T) {
	testScoreCount(t, 3, Cards{1, 2, 3})
}

package clean

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

func TestCleaner(t *testing.T) {
	playedCards := deck.Cards{2, 3, 4, 5, 6}
	testObject := NewCleaner(&playedCards)
	testObject.Clean()
	if len(playedCards) != 0 {
		t.Fatalf("Cards were not moved from played cards")
	}
}

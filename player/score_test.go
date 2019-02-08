package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/deck"
	"github.com/nikiforosFreespirit/msdb5/point"
)

func setupAndVerify(t *testing.T, cards deck.Cards, expectedScore uint8) {
	p := New()
	p.collect(cards)
	actualScore := point.Count(cards)
	if expectedScore != actualScore {
		t.Fatalf("Score should be %d but is %d", expectedScore, actualScore)
	}
}

func TestPlayerHasScoreOf0WhenCreated(t *testing.T) {
	setupAndVerify(t, deck.Cards{}, 0)
}

func TestPlayerHasScoreOf2WhenCollectingCard8AndOtherCardsScoring0(t *testing.T) {
	setupAndVerify(t, deck.Cards{8, 32, 16, 4, 35}, 2)
}

func TestPlayerHasScoreOf120WhenCollectingAllCards(t *testing.T) {
	setupAndVerify(t, deck.Deck(), 120)
}

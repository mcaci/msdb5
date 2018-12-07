package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card/set"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

func setupAndVerify(t *testing.T, cards set.Cards, expectedScore uint8) {
	p := New()
	p.Collect(cards)
	actualScore := p.Score(rule.Count)
	if expectedScore != actualScore {
		t.Fatalf("Score should be %d but is %d", expectedScore, actualScore)
	}
}

func TestPlayerHasScoreOf0WhenCreated(t *testing.T) {
	setupAndVerify(t, set.Cards{}, 0)
}

func TestPlayerHasScoreOf2WhenCollectingCard8AndOtherCardsScoring0(t *testing.T) {
	setupAndVerify(t, set.Cards{8, 32, 16, 4, 35}, 2)
}

func TestPlayerHasScoreOf120WhenCollectingAllCards(t *testing.T) {
	setupAndVerify(t, set.Deck(), 120)
}

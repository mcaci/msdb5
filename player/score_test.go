package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/rule"
)

func setupAndVerify(t *testing.T, cards card.Cards, expectedScore uint8) {
	p := New()
	p.collect(cards)
	actualScore := p.score(rule.Count)
	if expectedScore != actualScore {
		t.Fatalf("Score should be %d but is %d", expectedScore, actualScore)
	}
}

func TestPlayerHasScoreOf0WhenCreated(t *testing.T) {
	setupAndVerify(t, card.Cards{}, 0)
}

func TestPlayerHasScoreOf2WhenCollectingCard8AndOtherCardsScoring0(t *testing.T) {
	setupAndVerify(t, card.Cards{8, 32, 16, 4, 35}, 2)
}

func TestPlayerHasScoreOf120WhenCollectingAllCards(t *testing.T) {
	setupAndVerify(t, card.Deck(), 120)
}

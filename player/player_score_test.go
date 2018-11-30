package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/rule"

	"github.com/nikiforosFreespirit/msdb5/card/set"
)

func TestPlayerHasScoreOf0WhenCreated(t *testing.T) {
	p := New()
	actualScore := p.Score(rule.Count)
	if 0 != actualScore {
		t.Fatal("Score should be 0 at the beginning")
	}
}

func TestPlayerHasScoreOf2WhenCollectingCard8AndOtherCardsScoring0(t *testing.T) {
	p := New()
	p.Collect(set.Cards{8, 32, 16, 4, 35})
	expected := uint8(2)
	actual := p.Score(rule.Count)
	if expected != actual {
		t.Fatalf("Score should be %d but is %d", expected, actual)
	}
}

func TestPlayerHasScoreOf120WhenCollectingAllCards(t *testing.T) {
	p := New()
	p.Collect(set.Deck())
	expected := uint8(120)
	actual := p.Score(rule.Count)
	if expected != actual {
		t.Fatalf("Score should be %d but is %d", expected, actual)
	}
}

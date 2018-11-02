package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestPlayerHasNoCardsAtStartGame(t *testing.T) {
	p := New()
	collectedCards := p.Pile()
	if len(*collectedCards) > 0 {
		t.Fatal("Player should not have cards at creation")
	}
}

func TestPlayerWinsCards(t *testing.T) {
	p := New()
	cards := card.Cards{1, 2, 3, 4, 5}
	p.Pile().Add(cards...)

	collectedCards := p.Pile()
	if collectedCards == nil {
		t.Fatalf("Player should have %v but has %v", cards, collectedCards)
	}
}

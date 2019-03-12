package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

func TestPlayerHasNoCardsAtStartGame(t *testing.T) {
	if p := New(); len(p.pile) > 0 {
		t.Fatal("Player should not have cards at creation")
	}
}

func TestPlayerWinsCards(t *testing.T) {
	p := New()
	cards := deck.Cards{1, 2, 3, 4, 5}
	p.Collect(&cards)

	if p.pile == nil {
		t.Fatalf("Player should have %v but has %v", cards, p.pile)
	}
}

func TestPlayerCountPoints(t *testing.T) {
	p := New()
	cards := deck.Cards{1, 2, 3, 4, 5}
	p.Collect(&cards)

	if score := p.Count(func(card.ID) uint8 { return 1 }); score != 5 {
		t.Fatalf("Player should have 5 points but has %d", score)
	}
}

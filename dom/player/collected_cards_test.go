package player

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
)

func TestPlayerHasNoCardsAtStartGame(t *testing.T) {
	if p := New(); !IsHandEmpty(p) {
		t.Fatal("Player should not have cards at creation")
	}
}

func initPlayerForCollectTest() *Player {
	p := New()
	cards := deck.Cards{1, 2, 3, 4, 5}
	p.Collect(&cards)
	return p
}

func TestPlayerWinsCards(t *testing.T) {
	p := initPlayerForCollectTest()
	if p.pile == nil {
		t.Fatalf("Player should have %v but has %v", deck.Cards{1, 2, 3, 4, 5}, p.pile)
	}
}

func TestPlayerCountPoints(t *testing.T) {
	p := initPlayerForCollectTest()
	if score := p.Points(func(card.ID) uint8 { return 1 }); score != 5 {
		t.Fatalf("Player should have 5 points but has %d", score)
	}
}

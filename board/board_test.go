package board

import (
	"testing"
)

func TestBoardHasADeck(t *testing.T) {
	b := New()
	if b.Deck() == nil {
		t.Fatal("The board has no Deck")
	}
}

func TestBoardHas5Player(t *testing.T) {
	b := New()
	if b.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

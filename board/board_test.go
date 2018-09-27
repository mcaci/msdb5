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

func TestBoardsDeckReferenceIsTheSame(t *testing.T) {
	b := New()
	deck1 := b.Deck()
	deck2 := b.Deck()
	if deck1 != deck2 {
		t.Fatal("The deck is not the same each time is retrieved")
	}
}

func TestBoardHas5Player(t *testing.T) {
	b := New()
	if b.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

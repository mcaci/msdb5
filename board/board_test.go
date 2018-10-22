package board

import (
	"testing"
)

func TestBoardHasADeck(t *testing.T) {
	if b := New(); b.Deck() == nil {
		t.Fatal("The board has no Deck")
	}
}

func TestBoardHas5Player(t *testing.T) {
	if b := New(); b.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

func Test5PlayersDrawUntilDeckIsEmpty(t *testing.T) {
	if b := New(); len(b.Deck()) > 0 {
		t.Fatal("Not all cards have been distributed")
	}
}

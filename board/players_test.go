package board

import (
	"testing"
)

func TestBoardHas5Player(t *testing.T) {
	if b := New(); b.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

func TestPlayer1Has8Cards(t *testing.T) {
	if b := New(); len(*b.Players()[0].Hand()) != 8 {
		t.Fatalf("Player has %d cards", len(*b.Players()[0].Hand()))
	}
}

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

func TestPlayer1JoinsCheckName(t *testing.T) {
	b := New()
	if b.Join("Michi", "127.0.0.1"); b.Players()[0].Name() != "Michi" {
		t.Fatalf("Player's name was not registered correctly, found '%s'", b.Players()[0].Name())
	}
}

func TestPlayer1JoinsCheckIP(t *testing.T) {
	b := New()
	if b.Join("Michi", "127.0.0.1"); b.Players()[0].Host() != "127.0.0.1" {
		t.Fatalf("Player's ip was not registered correctly, found '%s'", b.Players()[0].Host())
	}
}

func TestPlayer2JoinsCheckName(t *testing.T) {
	b := New()
	b.Join("Michi", "127.0.0.1")
	if b.Join("Mary", "127.0.0.2"); b.Players()[1].Name() != "Mary" {
		t.Fatalf("Player's name was not registered correctly, found '%s'", b.Players()[1].Name())
	}
}

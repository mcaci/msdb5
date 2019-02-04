package board

import "testing"

func TestPlayer1JoinsCheckName(t *testing.T) {
	b := New()
	b.Join("Michi", "127.0.0.1")
	if name := b.Players()[0].Name(); name != "Michi" {
		t.Fatalf("Player's name was not registered correctly, found '%s'", name)
	}
}

func TestPlayer1JoinsCheckIP(t *testing.T) {
	b := New()
	b.Join("Michi", "127.0.0.1")
	if host := b.Players()[0].Host(); host != "127.0.0.1" {
		t.Fatalf("Player's ip was not registered correctly, found '%s'", host)
	}
}

func TestPlayer2JoinsCheckName(t *testing.T) {
	b := New()
	b.Join("Michi", "127.0.0.1")
	b.Join("Mary", "127.0.0.2")
	if name := b.Players()[1].Name(); name != "Mary" {
		t.Fatalf("Player's name was not registered correctly, found '%s'", name)
	}
}

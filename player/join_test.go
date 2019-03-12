package player

import "testing"

var p *Player

func init() {
	p := New()
	p.Join("Michi", "127.0.0.1")
}

func TestJoinPlayerSameToItself(t *testing.T) {
	if !p.IsSame(p) {
		t.Fatal("P should be equal to itself")
	}
}

func TestJoinPlayerName(t *testing.T) {
	if !p.IsName("Michi") {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerNameNotEmpty(t *testing.T) {
	if p.IsNameEmpty() {
		t.Fatal("Unexpected name being empty")
	}
}

func TestJoinPlayerHost(t *testing.T) {
	if !p.IsSameHost("127.0.0.1") {
		t.Fatal("Unexpected host")
	}
}

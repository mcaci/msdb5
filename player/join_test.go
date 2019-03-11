package player

import "testing"

func TestJoinPlayerName(t *testing.T) {
	p := New()
	p.Join("Michi", "127.0.0.1")
	if !p.IsName("Michi") {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerHost(t *testing.T) {
	p := New()
	p.Join("Michi", "127.0.0.1")
	if !p.IsSameHost("127.0.0.1") {
		t.Fatal("Unexpected host")
	}
}

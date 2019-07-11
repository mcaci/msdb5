package player

import "testing"

func initTest() *Player {
	p := New()
	p.RegisterAs("Michi")
	p.Join("127.0.0.1")
	return p
}

func TestJoinPlayerName(t *testing.T) {
	if p := initTest(); p.Name() != "Michi" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerNameNotEmpty(t *testing.T) {
	if p := initTest(); IsNameEmpty(p) {
		t.Fatal("Unexpected name being empty")
	}
}

func TestJoinPlayerHost(t *testing.T) {
	if p := initTest(); !p.IsSameHost("127.0.0.1") {
		t.Fatal("Unexpected host")
	}
}

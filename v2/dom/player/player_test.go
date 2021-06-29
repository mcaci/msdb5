package player

import (
	"testing"
)

func testP() *Player {
	p := New(&Options{For2P: true}).(*Player)
	p.RegisterAs("Me")
	return p
}

func TestNewPlayersAreNotSame(t *testing.T) {
	if Matching(testP())(testP()) {
		t.Fatal("Unexpected players being equal")
	}
}

func TestJoinPlayerName(t *testing.T) {
	if p := testP(); p.Name() != "Me" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := testP(); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}

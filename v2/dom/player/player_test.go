package player

import (
	"testing"
)

func initTest() *Player {
	p := New()
	p.RegisterAs("Me")
	return p
}

func TestNewPlayersAreNotSame(t *testing.T) {
	if Matching(initTest())(initTest()) {
		t.Fatal("Unexpected players being equal")
	}
}

func TestJoinPlayerName(t *testing.T) {
	if p := initTest(); p.Name() != "Me" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := initTest(); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}

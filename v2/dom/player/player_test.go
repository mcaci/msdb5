package player

import (
	"testing"
)

func testP(name string) Player {
	return New(&Options{For2P: true, Name: name})
}

func TestJoinPlayerName(t *testing.T) {
	if p := testP("Me"); p.Name() != "Me" {
		t.Fatal("Unexpected name")
	}
}

func TestJoinPlayerPileIsEmpty(t *testing.T) {
	if p := testP(""); len(*p.Pile()) != 0 {
		t.Fatal("Pile should be empty")
	}
}

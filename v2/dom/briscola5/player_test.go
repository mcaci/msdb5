package briscola5

import (
	"testing"
)

func testP(name string) Player {
	return *NewB5Player(name)
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

func TestPlayerHasFolded(t *testing.T) {
	p := NewB5Player("")
	if p.Fold(); !p.Folded() {
		t.Fatal("Player should have folded")
	}
}

func TestPlayerHasntFolded(t *testing.T) {
	p := NewB5Player("")
	if p.Folded() {
		t.Fatal("New player should not have folded")
	}
}

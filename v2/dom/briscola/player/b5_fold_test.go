package player

import (
	"testing"
)

func testB5P() *B5Player { return New(&Options{For5P: true}).(*B5Player) }

func TestPlayerHasFolded(t *testing.T) {
	p := testB5P()
	if p.Fold(); NotFolded(p) {
		t.Fatal("Player should have folded")
	}
}

func TestPlayerHasntFolded(t *testing.T) {
	p := testB5P()
	if Folded(p) {
		t.Fatal("New player should not have folded")
	}
}

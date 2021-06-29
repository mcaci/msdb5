package player

import (
	"testing"
)

func TestPlayerHasFolded(t *testing.T) {
	p := NewPlayer()
	if p.Fold(); !Folded(p) {
		t.Fatal("Player should have folded")
	}
}

func TestPlayerHasntFolded(t *testing.T) {
	p := NewPlayer()
	if Folded(p) {
		t.Fatal("Player should not have folded")
	}
}

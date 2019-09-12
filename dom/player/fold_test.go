package player

import (
	"testing"
)

func TestPlayerHasFolded(t *testing.T) {
	p := New()
	if p.Fold(); !Folded(p) {
		t.Fatal("Player should have folded")
	}
}

func TestPlayerHasntFolded(t *testing.T) {
	p := New()
	if Folded(p) {
		t.Fatal("Player should not have folded")
	}
}

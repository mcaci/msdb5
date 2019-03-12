package player

import (
	"testing"
)

func TestPlayerHasFolded(t *testing.T) {
	p := New()
	p.Fold()
	if p.NotFolded() {
		t.Fatal("Player should have folded")
	}
}

func TestPlayerHasntFolded(t *testing.T) {
	p := New()
	if p.Folded() {
		t.Fatal("Player should not have folded")
	}
}

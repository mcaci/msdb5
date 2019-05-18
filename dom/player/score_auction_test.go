package player

import (
	"testing"
)

func TestPlayerHasntFolded(t *testing.T) {
	p := New()
	if p.Folded() {
		t.Fatal("Player should not have folded")
	}
}

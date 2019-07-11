package player

import (
	"testing"
)

func TestPlayerHasntFolded(t *testing.T) {
	p := New()
	if Folded(p) {
		t.Fatal("Player should not have folded")
	}
}

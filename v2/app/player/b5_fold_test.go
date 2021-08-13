package player

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/briscola5"
)

func testB5P() *briscola5.Player { return New(&Options{For5P: true}).(*briscola5.Player) }

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

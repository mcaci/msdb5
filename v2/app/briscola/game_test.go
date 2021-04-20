package briscola

import (
	"testing"
)

func TestSideDeckProperty(t *testing.T) {
	g := NewGame(&Options{})
	if g.opts.WithName != "" {
		t.Errorf("error")
	}
	g = NewGame(&Options{WithName: "test"})
	if g.opts.WithName != "test" {
		t.Errorf("error")
	}
}

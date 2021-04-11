package game

import (
	"testing"
)

func TestSideDeckProperty(t *testing.T) {
	g := NewGame(&Options{})
	if g.opts.WithSide {
		t.Errorf("error")
	}
	g = NewGame(&Options{WithSide: true})
	if !g.opts.WithSide {
		t.Errorf("error")
	}
}

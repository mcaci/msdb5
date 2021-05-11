package briscola

import (
	"testing"

	"github.com/mcaci/msdb5/v2/dom/briscola"
)

func TestSideDeckProperty(t *testing.T) {
	g := NewGame(WithDefaultOptions)
	if g.opts.WithName != "" {
		t.Errorf("error")
	}
	g = NewGame(&Options{WithName: "test"})
	if g.opts.WithName != "test" {
		t.Errorf("error")
	}
}

func TestStartGameLenPlayers(t *testing.T) {
	g := NewGame(WithDefaultOptions)
	Start(g)
	if l := g.Players().Len(); l != 2 {
		t.Errorf("Expecting 2 players but got %d", l)
	}
}

func TestSetBriscola(t *testing.T) {
	g := NewGame(WithDefaultOptions)
	c := briscola.Card{Item: g.deck.Top()}
	Set(c, g)
	if ac := g.briscolaCard.ToID(); ac != c.ToID() {
		t.Errorf("Expecting card to be %v but was %v", c, ac)
	}
}

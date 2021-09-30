package briscola_test

import (
	"testing"

	"github.com/mcaci/msdb5/v3/briscola"
)

func TestSideDeckProperty(t *testing.T) {
	t.Parallel()
	g := briscola.NewGame(briscola.WithDefaultOptions)
	if g.Name != "" {
		t.Errorf("error")
	}
	g = briscola.NewGame(&briscola.Options{WithName: "test"})
	if g.Name != "test" {
		t.Errorf("error")
	}
}

func TestStartGameLenPlayersAndBriscola(t *testing.T) {
	t.Parallel()
	g := briscola.NewGame(briscola.WithDefaultOptions)
	if l := len(*g.Players()); l != 2 {
		t.Errorf("Expecting 2 players but got %d", l)
	}
}

func TestRegisterToGameAndStart(t *testing.T) {
	t.Parallel()
	g := briscola.NewGame(briscola.WithDefaultOptions)
	testSinglePlayerRegistering(t, "p1", 0, g)
	testSinglePlayerRegistering(t, "p2", 1, g)
	// game has started
	if ac := g.BriscolaCard(); ac == nil {
		t.Error("Expecting card to be something but was nil")
	}
}

func TestRegisterTooManyPlayers(t *testing.T) {
	t.Parallel()
	g := briscola.NewGame(briscola.WithDefaultOptions)
	testSinglePlayerRegistering(t, "p1", 0, g)
	testSinglePlayerRegistering(t, "p2", 1, g)
	// game has already started but let's try to register a new player
	err := briscola.Register("p3", g)
	if err == nil {
		t.Error("expecting an error here but got nothing")
	}
}

func testSinglePlayerRegistering(t *testing.T, name string, id int, g *briscola.Game) {
	err := briscola.Register(name, g)
	if err != nil {
		t.Errorf("Not expecting an error here but got %v", err)
	}
	if expected := (*g.Players())[id].Name(); expected != name {
		t.Errorf(`Expecting %q but got %q`, expected, name)
	}
}

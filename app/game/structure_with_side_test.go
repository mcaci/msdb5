package game

import (
	"testing"
)

func TestGameWithSideHas5Player(t *testing.T) {
	if gameTest := NewGame(true); gameTest.playersRef() == nil {
		t.Fatal("There are no Player")
	}
}

func TestGameWithSideHasNoPlayerInTurnAtStart(t *testing.T) {
	if gameTest := NewGame(true); gameTest.CurrentPlayer() == nil {
		t.Fatal("There are no Player in turn")
	}
}

func TestPlayer1Has7Cards(t *testing.T) {
	if gameTest := NewGame(true); len(*gameTest.players[0].Hand()) != 7 {
		t.Fatalf("Player has %d cards", len(*gameTest.players[0].Hand()))
	}
}

func TestPlayer4Has7Cards(t *testing.T) {
	if gameTest := NewGame(true); len(*gameTest.players[4].Hand()) != 7 {
		t.Fatalf("Player has %d cards", len(*gameTest.players[4].Hand()))
	}
}

func TestSideDeckHas5CardsWhenPresent(t *testing.T) {
	if gameTest := NewGame(true); len(gameTest.side) != 5 {
		t.Fatalf("Side deck has %d cards", len(gameTest.side))
	}
}

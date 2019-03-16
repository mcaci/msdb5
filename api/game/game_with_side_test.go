package game

import (
	"testing"
)

func TestGameWithSideHas5Player(t *testing.T) {
	if gameTest := NewGame(true); gameTest.Players() == nil {
		t.Fatal("The board has no Player")
	}
}

func TestGameWithSideHasNoPlayerInTurnAtStart(t *testing.T) {
	if gameTest := NewGame(true); gameTest.PlayerInTurn() == nil {
		t.Fatal("The board has no Player in turn")
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

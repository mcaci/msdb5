package game

import (
	"testing"
)

func TestGameWithSideHas5Player(t *testing.T) {
	gameTest := NewGame(true)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.players == nil {
		t.Fatal("There are no Player")
	}
}

func TestGameWithSideHasNoPlayerInTurnAtStart(t *testing.T) {
	gameTest := NewGame(true)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("There are no Player in turn")
	}
}

package orchestrator

import (
	"testing"
)

func TestPlayer1Joins(t *testing.T) {
	gameTest := NewGame()
	err := gameTest.Join("A", "127.0.0.1")
	if err != nil {
		t.Fatal("Single join operation was not successful")
	}
}

func TestPlayer1JoinsStatusIsJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "127.0.0.1")
	if gameTest.statusInfo != joining {
		t.Fatal("Status is not correct")
	}
}

func TestPlayer2Joins(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	err := gameTest.Join("Mary", "127.0.0.2")
	if err != nil {
		t.Fatal("Double join operation was not successful for second player")
	}
}

func TestPlayer2JoinsStatusIsJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	err := gameTest.Join("Mary", "127.0.0.2")
	if err != nil {
		t.Fatal("Double join operation was not successful for second player")
	}
}

func TestPlayer6CannotJoin(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	gameTest.Join("A", "127.0.0.3")
	gameTest.Join("gameTest", "127.0.0.4")
	gameTest.Join("C", "127.0.0.5")
	err := gameTest.Join("Nope", "127.0.0.6")
	if err == nil {
		t.Fatal("Player 'Nope' should not be joining as there is no sixth player")
	}
}

func TestPlayer6CannotJoinStatusChangesToAuction(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	gameTest.Join("A", "127.0.0.3")
	gameTest.Join("gameTest", "127.0.0.4")
	gameTest.Join("C", "127.0.0.5")
	err := gameTest.Join("Nope", "127.0.0.6")
	if err == nil {
		t.Fatal("Player 'Nope' should not be joining as there is no sixth player")
	}
}

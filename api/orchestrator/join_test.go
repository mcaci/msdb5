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

func TestPlayer1JoinsPhaseIsJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "127.0.0.1")
	if gameTest.phase != joining {
		t.Fatal("Phase is not correct")
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

func TestPlayer2JoinsPhaseIsJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	if gameTest.phase != joining {
		t.Fatal("Phase is not correct")
	}
}

func TestPlayer5Joins(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	gameTest.Join("A", "127.0.0.3")
	gameTest.Join("gameTester", "127.0.0.4")
	err := gameTest.Join("C", "127.0.0.5")
	if err != nil {
		t.Fatal("All players should have joined correctly")
	}
}

func TestPlayer5JoinsAndPhaseChangesToAuction(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	gameTest.Join("A", "127.0.0.3")
	gameTest.Join("gameTester", "127.0.0.4")
	gameTest.Join("C", "127.0.0.5")
	if gameTest.phase != scoreAuction {
		t.Fatal("Phase is not correct")
	}
}

func TestPlayer5JoinsAndSetPlayerInTurnToFirstJoiner(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	gameTest.Join("A", "127.0.0.3")
	gameTest.Join("gameTester", "127.0.0.4")
	gameTest.Join("C", "127.0.0.5")
	if gameTest.playerInTurn != gameTest.players[0] {
		t.Fatal("Player in turn is not set correctly")
	}
}

func TestPlayer6CannotJoin(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("Michi", "127.0.0.1")
	gameTest.Join("Mary", "127.0.0.2")
	gameTest.Join("A", "127.0.0.3")
	gameTest.Join("gameTester", "127.0.0.4")
	gameTest.Join("C", "127.0.0.5")
	err := gameTest.Join("Nope", "127.0.0.6")
	if err == nil {
		t.Fatal("Player 'Nope' should not be joining as there is no sixth player")
	}
}

func TestPlayerCannotJoinIfPhaseIsNotJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.phase = scoreAuction
	err := gameTest.Join("A", "127.0.0.1")
	if err == nil {
		t.Fatal("Player cannot join if phase is not joining")
	}
}

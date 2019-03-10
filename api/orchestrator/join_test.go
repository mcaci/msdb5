package orchestrator

import (
	"testing"
)

func TestPlayer1Joins(t *testing.T) {
	gameTest := NewGame()
	_, _, err := gameTest.join("Join#A", "127.0.0.101")
	if err != nil {
		t.Fatal("Single join operation was not successful")
	}
}

func TestPlayer1JoinsPhaseIsJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#A", "127.0.0.101")
	if gameTest.phase != joining {
		t.Fatal("Phase is not correct")
	}
}

func TestPlayer2Joins(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#Michi", "127.0.0.101")
	_, _, err := gameTest.join("Join#Mary", "127.0.0.102")
	if err != nil {
		t.Fatal("Double join operation was not successful for second player")
	}
}

func TestPlayer2JoinsPhaseIsJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#Michi", "127.0.0.101")
	gameTest.join("Join#Mary", "127.0.0.102")
	if gameTest.phase != joining {
		t.Fatal("Phase is not correct")
	}
}

func TestPlayer5Joins(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#Michi", "127.0.0.101")
	gameTest.join("Join#Mary", "127.0.0.102")
	gameTest.join("Join#A", "127.0.0.103")
	gameTest.join("Join#gameTester", "127.0.0.104")
	_, _, err := gameTest.join("Join#C", "127.0.0.105")
	if err != nil {
		t.Fatal("All players should have joined correctly")
	}
}

func TestPlayer5JoinsAndPhaseChangesToAuction(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#Michi", "127.0.0.101")
	gameTest.join("Join#Mary", "127.0.0.102")
	gameTest.join("Join#A", "127.0.0.103")
	gameTest.join("Join#gameTester", "127.0.0.104")
	gameTest.join("Join#C", "127.0.0.105")
	if gameTest.phase != scoreAuction {
		t.Fatal("Phase is not correct")
	}
}

func TestPlayer5JoinsAndSetPlayerInTurnToFirstJoiner(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#Michi", "127.0.0.101")
	gameTest.join("Join#Mary", "127.0.0.102")
	gameTest.join("Join#A", "127.0.0.103")
	gameTest.join("Join#gameTester", "127.0.0.104")
	gameTest.join("Join#C", "127.0.0.105")
	if gameTest.playerInTurn != 0 {
		t.Fatal("Player in turn is not set correctly")
	}
}

func TestPlayer6CannotJoin(t *testing.T) {
	gameTest := NewGame()
	gameTest.join("Join#Michi", "127.0.0.101")
	gameTest.join("Join#Mary", "127.0.0.102")
	gameTest.join("Join#A", "127.0.0.103")
	gameTest.join("Join#gameTester", "127.0.0.104")
	gameTest.join("Join#C", "127.0.0.105")
	_, _, err := gameTest.join("Join#Nope", "127.0.0.106")
	if err == nil {
		t.Fatal("Player 'Nope' should not be joining as there is no sixth player")
	}
}

func TestPlayerCannotJoinIfPhaseIsNotJoining(t *testing.T) {
	gameTest := NewGame()
	gameTest.phase = scoreAuction
	_, _, err := gameTest.join("Join#A", "127.0.0.101")
	if err == nil {
		t.Fatal("Player cannot join if phase is not joining")
	}
}

package orchestrator

import "testing"

func TestPlayerCannotChooseCompanionIfPhaseIsNotCompanion(t *testing.T) {
	gameTest := NewGame()
	err := gameTest.Nominate("3", "Cup", "100.1.1.1")
	if err == nil {
		t.Fatal("Nominate action not expected at beginning of game")
	}
}

func TestPlayerInTurnCanNominate(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.phase = companionChoice
	gameTest.playerInTurn = 0
	err := gameTest.Nominate("2", "Cudgel", "100.1.1.1")
	if err != nil {
		t.Fatal("Expecting in turn player to nominate companion card with success")
	}
}

func TestAnyOtherPlayerNotInTurnCantNominate(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.phase = companionChoice
	gameTest.playerInTurn = 1
	err := gameTest.Nominate("2", "Cudgel", "100.1.1.1")
	if err == nil {
		t.Fatal("Expecting not in turn player to being able to nominate companion card")
	}
}

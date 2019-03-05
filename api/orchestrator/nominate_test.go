package orchestrator

import "testing"

func TestPlayerCannotChooseCompanionIfPhaseIsNotCompanion(t *testing.T) {
	gameTest := NewGame()
	err := gameTest.Nominate("3", "Cup", "100.1.1.1")
	if err == nil {
		t.Fatal("Nominate action not expected at beginning of game")
	}
}

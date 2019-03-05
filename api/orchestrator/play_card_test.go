package orchestrator

import "testing"

func TestPlayerCannotPlayCardsIfPhaseIsNotPlay(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	err := gameTest.Play("3", "Sword", "100.1.1.1")
	if err == nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

func TestPlaysOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	gameTest.phase = playBriscola
	err := gameTest.Play("3", "Sword", "100.1.1.1")
	if err != nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

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

func TestPlayerPlaysOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	gameTest.phase = playBriscola
	err := gameTest.Play("3", "Sword", "100.1.1.1")
	if err != nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

func TestSecondPlayerCannotPlayCardBeforeFirstPlayer(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.phase = playBriscola
	gameTest.players[0].Hand().Add(34)
	gameTest.players[1].Hand().Add(15)
	err := gameTest.Play("5", "Cup", "100.1.1.2")
	if err == nil {
		t.Fatal("Expecting error for second player not being able to play cards before first player has")
	}
}

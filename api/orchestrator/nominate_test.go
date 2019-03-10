package orchestrator

import "testing"

func TestPlayerCannotChooseCompanionIfPhaseIsNotCompanion(t *testing.T) {
	gameTest := NewGame()
	_, _, err := gameTest.nominate("Companion#3#Cup", "127.0.0.31")
	if err == nil {
		t.Fatal("Nominate action not expected at beginning of game")
	}
}

func TestPlayerInTurnCanNominate(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.31")
	gameTest.phase = companionChoice
	gameTest.playerInTurn = 0
	_, _, err := gameTest.nominate("Companion#2#Cudgel", "127.0.0.31")
	if err != nil {
		t.Fatal("Expecting in turn player to nominate companion card with success")
	}
}

func TestNominatedInfoIsFilled(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.31")
	gameTest.phase = companionChoice
	gameTest.playerInTurn = 0
	gameTest.nominate("Companion#2#Cudgel", "127.0.0.31")
	if gameTest.companion.Card() != 32 {
		t.Fatal("Expecting in turn player to nominate companion card with success")
	}
}

func TestAnyOtherPlayerNotInTurnCantNominate(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.31")
	gameTest.players[1].Join("B", "127.0.0.32")
	gameTest.phase = companionChoice
	gameTest.playerInTurn = 1
	_, _, err := gameTest.nominate("Companion#2#Cudgel", "127.0.0.31")
	if err == nil {
		t.Fatal("Expecting not in turn player to being able to nominate companion card")
	}
}

func TestTransitionToPlayPhase(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.31")
	gameTest.phase = companionChoice
	gameTest.playerInTurn = 0
	gameTest.nominate("Companion#2#Cudgel", "127.0.0.31")
	if gameTest.phase != playBriscola {
		t.Fatal("Expecting in turn player to nominate companion card with success")
	}
}

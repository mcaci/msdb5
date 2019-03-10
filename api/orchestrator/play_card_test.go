package orchestrator

import (
	"testing"
)

func TestPlayerCannotPlayCardsIfPhaseIsNotPlay(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	err := gameTest.Play("3", "Sword", "100.1.1.1")
	if err == nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

func TestPlayerPlaysOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	gameTest.phase = playBriscola
	err := gameTest.Play("3", "Sword", "100.1.1.1")
	if err != nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

func TestNothingHappensIfPlayerPlaysNotOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	gameTest.phase = playBriscola
	err := gameTest.Play("0", "Coin", "100.1.1.1")
	if err == nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

func TestSceondPlayerCannotPlayIfFirstPlayerPlaysNotOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "100.1.1.1")
	gameTest.players[1].Join("B", "100.1.1.2")
	gameTest.players[1].Hand().Add(1)
	gameTest.phase = playBriscola
	gameTest.Play("0", "Coin", "100.1.1.1")
	err := gameTest.Play("1", "Coin", "100.1.1.2")
	if err == nil {
		t.Fatal("Second player should be waiting for first one to play proper card")
	}
}

func TestSecondPlayerCannotPlayCardBeforeFirstPlayer(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "100.1.1.1")
	gameTest.players[1].Join("B", "100.1.1.2")
	gameTest.phase = playBriscola
	gameTest.players[0].Hand().Add(34)
	gameTest.players[1].Hand().Add(15)
	err := gameTest.Play("5", "Cup", "100.1.1.2")
	if err == nil {
		t.Fatal("Expecting error for second player not being able to play cards before first player has")
	}
}

func TestSecondPlayerCanPlayCardAfterFirstPlayer(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "100.1.1.1")
	gameTest.players[1].Join("B", "100.1.1.2")
	gameTest.phase = playBriscola
	gameTest.players[0].Hand().Add(34)
	gameTest.players[1].Hand().Add(15)
	gameTest.Play("4", "Cudgel", "100.1.1.1")
	err := gameTest.Play("5", "Cup", "100.1.1.2")
	if err != nil {
		t.Fatal("Expecting second player being able to play cards after first player has")
	}
}

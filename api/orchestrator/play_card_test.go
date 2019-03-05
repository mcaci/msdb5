package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

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

func TestNothingHappensIfPlayerPlaysNotOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.players[0].Hand().Add(23)
	gameTest.phase = playBriscola
	err := gameTest.Play("0", "Coin", "100.1.1.1")
	if err == nil {
		t.Fatal("Play card action not expected at beginning of game")
	}
}

func TestSceondPlayerCannotPlayIfFirstPlayerPlaysNotOwnedCard(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.players[1].Hand().Add(1)
	gameTest.phase = playBriscola
	gameTest.Play("0", "Coin", "100.1.1.1")
	err := gameTest.Play("1", "Coin", "100.1.1.2")
	if err == nil {
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

func TestSecondPlayerCanPlayCardAfterFirstPlayer(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.phase = playBriscola
	gameTest.players[0].Hand().Add(34)
	gameTest.players[1].Hand().Add(15)
	gameTest.Play("4", "Cudgel", "100.1.1.1")
	err := gameTest.Play("5", "Cup", "100.1.1.2")
	if err != nil {
		t.Fatal("Expecting second player being able to play cards after first player has")
	}
}

func TestCompleteRound(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.Join("C", "100.1.1.3")
	gameTest.Join("D", "100.1.1.4")
	gameTest.Join("E", "100.1.1.5")
	gameTest.phase = playBriscola
	gameTest.companion = *companion.New(card.ID(9), gameTest.players[2])
	pile := deck.Cards{}
	for i, pl := range gameTest.players {
		pl.Hand().Move(&pile)
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	gameTest.Play("5", "Coin", "100.1.1.1")
	gameTest.Play("7", "Coin", "100.1.1.2")
	gameTest.Play("9", "Coin", "100.1.1.3")
	gameTest.Play("1", "Cup", "100.1.1.4")
	err := gameTest.Play("3", "Cup", "100.1.1.5")
	if err != nil {
		t.Fatal("Expecting full round to happen")
	}
}

func TestCompletedGame(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.Join("C", "100.1.1.3")
	gameTest.Join("D", "100.1.1.4")
	gameTest.Join("E", "100.1.1.5")
	gameTest.phase = playBriscola
	gameTest.companion = *companion.New(card.ID(9), gameTest.players[2])
	pile := deck.Cards{}
	for i, pl := range gameTest.players {
		pl.Hand().Move(&pile)
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	pile = nil
	gameTest.Play("5", "Coin", "100.1.1.1")
	gameTest.Play("7", "Coin", "100.1.1.2")
	gameTest.Play("9", "Coin", "100.1.1.3")
	gameTest.Play("1", "Cup", "100.1.1.4")
	gameTest.Play("3", "Cup", "100.1.1.5")
	if gameTest.phase != end {
		t.Fatal("Expecting transition to end game and scoring")
	}
}

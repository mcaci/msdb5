package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
)

func mockGameTest() (*Game, error) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.Join("C", "100.1.1.3")
	gameTest.Join("D", "100.1.1.4")
	gameTest.Join("E", "100.1.1.5")
	gameTest.companion = *companion.New(card.ID(9), gameTest.players[2])
	gameTest.phase = playBriscola
	fold4Players(gameTest)
	gameTest.Play("5", "Coin", "100.1.1.1")
	gameTest.Play("7", "Coin", "100.1.1.2")
	gameTest.Play("9", "Coin", "100.1.1.3")
	gameTest.Play("1", "Cup", "100.1.1.4")
	err := gameTest.Play("3", "Cup", "100.1.1.5")
	return gameTest, err
}

func fold4Players(gameTest *Game) {
	for i, pl := range gameTest.players {
		pl.Hand().Clear()
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
}

func TestCompleteRoundWithNoError(t *testing.T) {
	_, err := mockGameTest()
	if err != nil {
		t.Fatal("Expecting full round to happen")
	}
}

func TestCompleteRoundWithNextPlayerSelection(t *testing.T) {
	gameTest, _ := mockGameTest()
	if gameTest.playerInTurn != 2 {
		t.Fatalf("C should be next player to start, but is %d", gameTest.playerInTurn)
	}
}

func TestCompletedGame(t *testing.T) {
	gameTest, _ := mockGameTest()
	if gameTest.phase != end {
		t.Fatal("Expecting transition to end game and scoring")
	}
}

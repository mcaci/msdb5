package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
)

func mockGameTest() (*Game, error) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.41")
	gameTest.players[1].Join("B", "127.0.0.42")
	gameTest.players[2].Join("C", "127.0.0.43")
	gameTest.players[3].Join("D", "127.0.0.44")
	gameTest.players[4].Join("E", "127.0.0.45")
	gameTest.companion = *companion.New(card.ID(9), gameTest.players[2])
	gameTest.phase = playBriscola
	fold4Players(gameTest)
	gameTest.play("Card#5#Coin", "127.0.0.41")
	gameTest.play("Card#7#Coin", "127.0.0.42")
	gameTest.play("Card#9#Coin", "127.0.0.43")
	gameTest.play("Card#1#Cup", "127.0.0.44")
	_, _, err := gameTest.play("Card#3#Cup", "127.0.0.45")
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

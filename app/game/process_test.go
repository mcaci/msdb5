package game

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/app/game/start"
	"github.com/mcaci/msdb5/dom/phase"
)

func fakeGameSetup(withSide bool) *Game {
	gameTest := fakeGame(withSide)
	messageBufferSize := 256
	playerChannel := make(chan []byte, messageBufferSize)
	start.Join(gameTest, "127.0.0.51", playerChannel)
	start.Join(gameTest, "127.0.0.52", playerChannel)
	start.Join(gameTest, "127.0.0.53", playerChannel)
	start.Join(gameTest, "127.0.0.54", playerChannel)
	start.Join(gameTest, "127.0.0.55", playerChannel)
	if withSide {
		gameTest.side.Clear()
		gameTest.side.Add(*card.MustID(31))
	}
	for i, pl := range gameTest.players {
		pl.Hand().Clear()
		pl.Hand().Add(*card.MustID(uint8(2*i + 5)))
		if i > 1 {
			pl.Fold()
		}
	}
	return gameTest
}

func fakeGamePlay(gameTest *Game) {
	gameTest.Process("Join#A", "127.0.0.51")
	gameTest.Process("Join#B", "127.0.0.52")
	gameTest.Process("Join#C", "127.0.0.53")
	gameTest.Process("Join#D", "127.0.0.54")
	gameTest.Process("Join#E", "127.0.0.55")
	gameTest.Process("Auction#80", "127.0.0.51")
	gameTest.Process("Auction#79", "127.0.0.52")
	if gameTest.IsSideUsed() {
		gameTest.Process("Exchange#5#Coin", "127.0.0.51")
		gameTest.Process("Exchange#0#Coin", "127.0.0.51")
	}
	gameTest.Process("Companion#7#Coin", "127.0.0.51")

	if gameTest.IsSideUsed() {
		gameTest.Process("Card#1#Cudgel", "127.0.0.51")
	} else {
		gameTest.Process("Card#5#Coin", "127.0.0.51")
	}
	gameTest.Process("Card#7#Coin", "127.0.0.52")
	gameTest.Process("Card#9#Coin", "127.0.0.53")
	gameTest.Process("Card#1#Cup", "127.0.0.54")
	gameTest.Process("Card#3#Cup", "127.0.0.55")
}

func TestCompletedGameReturningScoreInfoWithSide(t *testing.T) {
	gameTest := fakeGameSetup(true)
	fakeGamePlay(gameTest)
	if gameTest.phase != phase.End {
		t.Fatalf("Expecting transition to end game and scoring but current phase was: %s", gameTest.phase)
	}
}

func TestCompletedGameHasSideDeckEmpty(t *testing.T) {
	gameTest := fakeGameSetup(true)
	fakeGamePlay(gameTest)
	if len(gameTest.side) != 0 {
		t.Fatalf("Expecting side deck to be empty but is: %v", gameTest.side)
	}
}

func TestCompletedGameReturningScoreInfoWithNoSide(t *testing.T) {
	gameTest := fakeGameSetup(false)
	fakeGamePlay(gameTest)
	if gameTest.phase != phase.End {
		t.Fatalf("Expecting transition to end game and scoring but current phase was: %s", gameTest.phase)
	}
}

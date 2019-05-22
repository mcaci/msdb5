package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

func TestCompletedGameReturningScoreInfo(t *testing.T) {
	gameTest := NewAction(false)
	gameTest.Process("Join#A", "127.0.0.51")
	gameTest.Process("Join#B", "127.0.0.52")
	gameTest.Process("Join#C", "127.0.0.53")
	gameTest.Process("Join#D", "127.0.0.54")
	gameTest.Process("Join#E", "127.0.0.55")
	game := gameTest.(*Game)
	game.SetCompanion(card.ID(9), game.Players()[2])
	for i, pl := range game.Players() {
		pl.Hand().Clear()
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	game.phase = 4
	gameTest.Process("Card#5#Coin", "127.0.0.51")
	gameTest.Process("Card#7#Coin", "127.0.0.52")
	gameTest.Process("Card#9#Coin", "127.0.0.53")
	gameTest.Process("Card#1#Cup", "127.0.0.54")
	info := gameTest.Process("Card#3#Cup", "127.0.0.55")
	if info.ForAll() == "" {
		t.Log(info.Err())
		t.Fatal("Expecting transition to end game and scoring")
	}
}

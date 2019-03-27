package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/game"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

func TestCompletedGameReturningScoreInfo(t *testing.T) {
	gameTest := NewAction(false)
	gameTest.Action("Join#A", "127.0.0.51")
	gameTest.Action("Join#B", "127.0.0.52")
	gameTest.Action("Join#C", "127.0.0.53")
	gameTest.Action("Join#D", "127.0.0.54")
	gameTest.Action("Join#E", "127.0.0.55")
	o := gameTest.(*Orchestrator)
	o.game.SetCompanion(card.ID(9), o.game.Players()[2])
	for i, pl := range o.game.Players() {
		pl.Hand().Clear()
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	o.game.NextPhase(game.PlayingCards)
	gameTest.Action("Card#5#Coin", "127.0.0.51")
	gameTest.Action("Card#7#Coin", "127.0.0.52")
	gameTest.Action("Card#9#Coin", "127.0.0.53")
	gameTest.Action("Card#1#Cup", "127.0.0.54")
	scoreInfo, _, _ := gameTest.Action("Card#3#Cup", "127.0.0.55")
	if scoreInfo == "" {
		t.Fatal("Expecting transition to end game and scoring")
	}
}

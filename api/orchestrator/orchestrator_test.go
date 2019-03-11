package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
)

func TestActionCreationAndAuctionUsage(t *testing.T) {
	gameTest := NewAction()
	gameTest.Action("Join#A", "127.0.0.51")
	_, pInfo, _ := gameTest.Action("Auction#102", "127.0.0.51")
	if pInfo == nil {
		t.Fatal("Auction action was not properly performed")
	}
}

func TestActionCreationAndPCompanionUsage(t *testing.T) {
	gameTest := NewAction()
	gameTest.Action("Join#A", "127.0.0.51")
	_, pInfo, _ := gameTest.Action("Companion#3#Cup", "127.0.0.51")
	if pInfo == nil {
		t.Fatal("Companion action was not properly performed")
	}
}

func TestActionCreationAndPlayCardUsage(t *testing.T) {
	gameTest := NewAction()
	gameTest.Action("Join#A", "127.0.0.51")
	_, pInfo, _ := gameTest.Action("Card#6#Cudgel", "127.0.0.51")
	if pInfo == nil {
		t.Fatal("Card action was not properly performed")
	}
}

func TestCompletedGameReturningScoreInfo(t *testing.T) {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.51")
	gameTest.players[1].Join("B", "127.0.0.52")
	gameTest.players[2].Join("C", "127.0.0.53")
	gameTest.players[3].Join("D", "127.0.0.54")
	gameTest.players[4].Join("E", "127.0.0.55")
	gameTest.phase = playBriscola
	gameTest.companion = *companion.New(card.ID(9), gameTest.players[2])
	for i, pl := range gameTest.players {
		pl.Hand().Clear()
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	gameTest.Action("Card#5#Coin", "127.0.0.51")
	gameTest.Action("Card#7#Coin", "127.0.0.52")
	gameTest.Action("Card#9#Coin", "127.0.0.53")
	gameTest.Action("Card#1#Cup", "127.0.0.54")
	scoreInfo, _, _ := gameTest.Action("Card#3#Cup", "127.0.0.55")
	if scoreInfo == nil {
		t.Log(scoreInfo)
		t.Fatal("Expecting transition to end game and scoring")
	}
}

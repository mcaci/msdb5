package orchestrator

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/companion"
)

func TestActionCreationAndAuctionUsage(t *testing.T) {
	gameTest := NewAction()
	gameTest.Action("Join#A", "100.1.1.1")
	_, pInfo, _ := gameTest.Action("Auction#102", "100.1.1.1")
	if pInfo == nil {
		t.Fatal("Auction action was not properly performed")
	}
}

func TestActionCreationAndPCompanionUsage(t *testing.T) {
	gameTest := NewAction()
	gameTest.Action("Join#A", "100.1.1.1")
	_, pInfo, _ := gameTest.Action("Companion#3#Cup", "100.1.1.1")
	if pInfo == nil {
		t.Fatal("Companion action was not properly performed")
	}
}

func TestActionCreationAndPlayCardUsage(t *testing.T) {
	gameTest := NewAction()
	gameTest.Action("Join#A", "100.1.1.1")
	_, pInfo, _ := gameTest.Action("Card#6#Cudgel", "100.1.1.1")
	if pInfo == nil {
		t.Fatal("Card action was not properly performed")
	}
}

func TestCompletedGameReturningScoreInfo(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.Join("C", "100.1.1.3")
	gameTest.Join("D", "100.1.1.4")
	gameTest.Join("E", "100.1.1.5")
	gameTest.phase = playBriscola
	gameTest.companion = *companion.New(card.ID(9), gameTest.players[2])
	for i, pl := range gameTest.players {
		pl.Hand().Clear()
		pl.Hand().Add(card.ID(2*i + 5))
		if i > 0 {
			pl.Fold()
		}
	}
	gameTest.Play("5", "Coin", "100.1.1.1")
	gameTest.Play("7", "Coin", "100.1.1.2")
	gameTest.Play("9", "Coin", "100.1.1.3")
	gameTest.Play("1", "Cup", "100.1.1.4")
	scoreInfo, _, _ := gameTest.Action("Card#3#Cup", "100.1.1.5")
	if scoreInfo == nil {
		t.Log(scoreInfo)
		t.Fatal("Expecting transition to end game and scoring")
	}
}

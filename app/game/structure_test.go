package game

import (
	"testing"
)

func TestGameSetsFirstPlayerAsCurrent(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("Current player should be the first player")
	}
}

func TestSideDeckHasNoCardsWhenAbsent(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.IsSideUsed() {
		t.Fatalf("Side deck has %d cards", len(gameTest.side))
	}
}

func TestPlayedCardsAreNotPresentAtCreation(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.cardsOnTheBoard() != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.cardsOnTheBoard())
	}
}

func TestAuctionScoreIsZeroAtCreation(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.auctionScore != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.auctionScore)
	}
}

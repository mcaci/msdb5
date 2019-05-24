package game

import (
	"testing"
)

func TestGameHas5Player(t *testing.T) {
	if gameTest := NewGame(false); gameTest.players == nil {
		t.Fatal("There are no Player")
	}
}

func TestGameSetsFirstPlayerAsCurrent(t *testing.T) {
	if gameTest := NewGame(false); gameTest.CurrentPlayer() == nil {
		t.Fatal("Current player should be the first player")
	}
}

func TestPlayer1Has8Cards(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.players[0].Hand()) != 8 {
		t.Fatalf("Player has %d cards", len(*gameTest.players[0].Hand()))
	}
}

func TestSideDeckHasNoCardsWhenAbsent(t *testing.T) {
	if gameTest := NewGame(false); gameTest.IsSideUsed() {
		t.Fatalf("Side deck has %d cards", len(gameTest.side))
	}
}

func TestPlayedCardsAreNotPresentAtCreation(t *testing.T) {
	if gameTest := NewGame(false); gameTest.cardsOnTheBoard() != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.cardsOnTheBoard())
	}
}

func TestAuctionScoreIsZeroAtCreation(t *testing.T) {
	if gameTest := NewGame(false); gameTest.auctionScore != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.auctionScore)
	}
}

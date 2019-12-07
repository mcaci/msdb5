package game

import (
	"testing"

	"github.com/mcaci/msdb5/app/game/start"
)

func fakeGame(hasSide bool) *Game {
	return NewGame(hasSide)
}

func TestGameSetsFirstPlayerAsCurrent(t *testing.T) {
	gameTest := fakeGame(false)
	start.Join(gameTest, "127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("Current player should be the first player")
	}
}

func TestSideDeckHasNoCardsWhenAbsent(t *testing.T) {
	gameTest := fakeGame(false)
	start.Join(gameTest, "127.0.0.51", make(chan []byte))
	if gameTest.IsSideUsed() {
		t.Fatalf("Side deck has %d cards", len(gameTest.side))
	}
}

func TestPlayedCardsAreNotPresentAtCreation(t *testing.T) {
	gameTest := fakeGame(false)
	start.Join(gameTest, "127.0.0.51", make(chan []byte))
	if len(gameTest.playedCards) == 5 {
		t.Fatal("Side deck is expected to have no more than 5 cards")
	}
}

func TestAuctionScoreIsZeroAtCreation(t *testing.T) {
	gameTest := fakeGame(false)
	start.Join(gameTest, "127.0.0.51", make(chan []byte))
	if gameTest.auctionScore != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.auctionScore)
	}
}

func TestGameWithSideHas5Player(t *testing.T) {
	gameTest := fakeGame(true)
	start.Join(gameTest, "127.0.0.51", make(chan []byte))
	if gameTest.players == nil {
		t.Fatal("There are no Player")
	}
}

func TestGameWithSideHasNoPlayerInTurnAtStart(t *testing.T) {
	gameTest := fakeGame(true)
	start.Join(gameTest, "127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("There are no Player in turn")
	}
}

func TestGameWithSideAndAuctionOf101Shows2SideDeckCards(t *testing.T) {
	gameTest := fakeGame(true)
	gameTest.SetAuction(101)
	if len(gameTest.SideSubset()) != 2 {
		t.Fatal("Expecting two cards to show")
	}
}

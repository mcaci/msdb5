package game

import (
	"testing"

	"golang.org/x/text/language"
)

func fakeGame(hasSide bool) *Game {
	return NewGame(hasSide, language.English)
}

func TestRegisterPlayerHasLocalhostOrigin(t *testing.T) {
	testGame := fakeGame(false)
	playerInfo := "localhost"
	testGame.Join(playerInfo, make(chan []byte))
	if p := testGame.players[0]; p == nil {
		t.Fatalf("Player %s is expected to exist", playerInfo)
	}
}

func TestGameSetsFirstPlayerAsCurrent(t *testing.T) {
	gameTest := fakeGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("Current player should be the first player")
	}
}

func TestSideDeckHasNoCardsWhenAbsent(t *testing.T) {
	gameTest := fakeGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.IsSideUsed() {
		t.Fatalf("Side deck has %d cards", len(gameTest.side))
	}
}

func TestPlayedCardsAreNotPresentAtCreation(t *testing.T) {
	gameTest := fakeGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if len(gameTest.playedCards) == 5 {
		t.Fatal("Side deck is expected to have no more than 5 cards")
	}
}

func TestAuctionScoreIsZeroAtCreation(t *testing.T) {
	gameTest := fakeGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.auctionScore != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.auctionScore)
	}
}

func TestGameWithSideHas5Player(t *testing.T) {
	gameTest := fakeGame(true)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.players == nil {
		t.Fatal("There are no Player")
	}
}

func TestGameWithSideHasNoPlayerInTurnAtStart(t *testing.T) {
	gameTest := fakeGame(true)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("There are no Player in turn")
	}
}

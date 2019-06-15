package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func TestRegisterPlayerHasLocalhostOrigin(t *testing.T) {
	testGame := NewGame(false)
	playerInfo := "localhost"
	testGame.Join(playerInfo, make(chan []byte))
	testPlayers := testGame.players
	_, _, err := testPlayers.Find(func(p *player.Player) bool { return p.IsSameHost(playerInfo) })
	if err != nil {
		t.Fatalf("Player %s is expected to exist", playerInfo)
	}
}

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
	if gameTest.CardsOnTheBoard() != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.CardsOnTheBoard())
	}
}

func TestAuctionScoreIsZeroAtCreation(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.auctionScore != 0 {
		t.Fatalf("Side deck has %d cards", gameTest.auctionScore)
	}
}

func TestGameWithSideHas5Player(t *testing.T) {
	gameTest := NewGame(true)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.players == nil {
		t.Fatal("There are no Player")
	}
}

func TestGameWithSideHasNoPlayerInTurnAtStart(t *testing.T) {
	gameTest := NewGame(true)
	gameTest.Join("127.0.0.51", make(chan []byte))
	if gameTest.CurrentPlayer() == nil {
		t.Fatal("There are no Player in turn")
	}
}

package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func TestGameHas5Player(t *testing.T) {
	if gameTest := NewGame(false); gameTest.Players() == nil {
		t.Fatal("There are no Player")
	}
}
func TestGameHasNoPlayerInTurnAtStart(t *testing.T) {
	if gameTest := NewGame(false); gameTest.PlayerInTurn() == nil {
		t.Fatal("There are no Player in turn")
	}
}

func TestPlayer1Has8Cards(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.players[0].Hand()) != 8 {
		t.Fatalf("Player has %d cards", len(*gameTest.players[0].Hand()))
	}
}

func TestSideDeckHasNoCardsWhenAbsent(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.SideDeck()) != 0 {
		t.Fatalf("Side deck has %d cards", len(*gameTest.SideDeck()))
	}
}

func TestPlayedCardsAreNotPresentAtCreation(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.PlayedCards()) != 0 {
		t.Fatalf("Side deck has %d cards", len(*gameTest.PlayedCards()))
	}
}

func TestAuctionScoreIsZeroAtCreation(t *testing.T) {
	if gameTest := NewGame(false); *gameTest.AuctionScore() != 0 {
		t.Fatalf("Side deck has %d cards", *gameTest.AuctionScore())
	}
}

func TestNextPhase(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.NextPhase(1)
	if gameTest.CurrentPhase() != 1 {
		t.Fatal("Current phase should be 1")
	}
}

func TestSetCompanionAndBriscolaSeed(t *testing.T) {
	testGame := NewGame(false)
	testGame.SetCompanion(1, player.New())
	if testGame.BriscolaSeed() != card.Coin {
		t.Fatal("Expecting coin as briscola")
	}
}

func TestSetCompanionAndPlayerReference(t *testing.T) {
	testGame := NewGame(false)
	testGame.SetCompanion(1, player.New())
	if testGame.Companion() == nil {
		t.Fatal("Companion to be set")
	}
}

func TestNextPlayer(t *testing.T) {
	testGame := NewGame(false)
	testGame.NextPlayer(func(uint8) uint8 { return 3 })
	if testGame.playerInTurn != 3 {
		t.Fatal("current player index should be 3")
	}
}

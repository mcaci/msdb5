package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/app/phase"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func TestGameHas5Player(t *testing.T) {
	if gameTest := NewGame(false); gameTest.playersRef() == nil {
		t.Fatal("There are no Player")
	}
}
func TestGameHasNoPlayerInTurnAtStart(t *testing.T) {
	if gameTest := NewGame(false); gameTest.CurrentPlayer() == nil {
		t.Fatal("There are no Player in turn")
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

func TestSetCompanionAndBriscolaSeed(t *testing.T) {
	testGame := NewGame(false)
	testGame.setCompanion(1, player.New())
	if testGame.briscola() != card.Coin {
		t.Fatal("Expecting coin as briscola")
	}
}

func TestSetCompanionAndPlayerReference(t *testing.T) {
	testGame := NewGame(false)
	testGame.setCompanion(1, player.New())
	if testGame.companion.Ref() == nil {
		t.Fatal("Companion to be set")
	}
}

func TestNextPlayer(t *testing.T) {
	testGame := NewGame(false)
	testGame.playerInTurn = nextPlayer(testGame, phase.Joining, 2)
	if testGame.playerInTurn != 3 {
		t.Fatal("current player index should be 3")
	}
}

func TestNextPlayerInsideAuction(t *testing.T) {
	testGame := NewGame(false)
	for i, player := range testGame.playersRef() {
		if i == 0 {
			continue
		}
		player.Fold()
	}
	testGame.playerInTurn = nextPlayer(testGame, phase.InsideAuction, 2)
	if testGame.playerInTurn != 0 {
		t.Fatal("current player index should be 0")
	}
}

func TestNextPlayerWhenExchangingCards(t *testing.T) {
	testGame := NewGame(false)
	testGame.playerInTurn = nextPlayer(testGame, phase.ExchangingCards, 2)
	if testGame.playerInTurn != 2 {
		t.Fatal("current player index should be 0")
	}
}

package game

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestBoardHas5Player(t *testing.T) {
	if gameTest := NewGame(false); gameTest.Players() == nil {
		t.Fatal("The board has no Player")
	}
}
func TestGameHasNoPlayerInTurnAtStart(t *testing.T) {
	if gameTest := NewGame(false); gameTest.PlayerInTurn() == nil {
		t.Fatal("The board has no Player in turn")
	}
}

func TestPlayer1Has8Cards(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.players[0].Hand()) != 8 {
		t.Fatalf("Player has %d cards", len(*gameTest.players[0].Hand()))
	}
}

func TestSideDeckHasNoCardsWhenAbsent(t *testing.T) {
	if gameTest := NewGame(false); len(gameTest.side) != 0 {
		t.Fatalf("Side deck has %d cards", len(gameTest.side))
	}
}

func TestBoardHasASetOfPlayedCards(t *testing.T) {
	if gameTest := NewGame(false); gameTest.board.PlayedCards() == nil {
		t.Fatal("The board has no set of played cards")
	}
}

func TestBoardsEmptySetOfPlayedCardsContainsNoCards(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.board.PlayedCards()) > 0 {
		t.Fatal("The deck should be empty at this point")
	}
}

func TestBoardAuctionScoreAtCreationIs0(t *testing.T) {
	if gameTest := NewGame(false); gameTest.Board().AuctionScore() != 0 {
		t.Fatalf("Auction score for a new board should be 0 but is %d", gameTest.board.AuctionScore())
	}
}

func TestNextPhase(t *testing.T) {
	gameTest := NewGame(false)
	gameTest.NextPhase()
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

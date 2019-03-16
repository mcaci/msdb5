package game

import (
	"testing"
)

func TestBoardHas5Player(t *testing.T) {
	if gameTest := NewGame(false); gameTest.players == nil {
		t.Fatal("The board has no Player")
	}
}

func TestPlayer1Has8Cards(t *testing.T) {
	if gameTest := NewGame(false); len(*gameTest.players[0].Hand()) != 8 {
		t.Fatalf("Player has %d cards", len(*gameTest.players[0].Hand()))
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
	if gameTest := NewGame(false); gameTest.board.AuctionScore() != 0 {
		t.Fatalf("Auction score for a new board should be 0 but is %d", gameTest.board.AuctionScore())
	}
}

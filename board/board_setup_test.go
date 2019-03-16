package board

import (
	"testing"
)

func TestBoardHasASetOfPlayedCards(t *testing.T) {
	if b := New(); b.PlayedCards() == nil {
		t.Fatal("The board has no set of played cards")
	}
}

func TestBoardsEmptySetOfPlayedCardsContainsNoCards(t *testing.T) {
	if b := New(); len(*b.PlayedCards()) > 0 {
		t.Fatal("The deck should be empty at this point")
	}
}

func TestBoardAuctionScoreAtCreationIs0(t *testing.T) {
	if b := New(); b.AuctionScore() != 0 {
		t.Fatalf("Auction score for a new board should be 0 but is %d", b.AuctionScore())
	}
}

func TestBoardSetAuctionScoreTo70(t *testing.T) {
	b := New()
	b.SetAuctionScore(71)
	if b.AuctionScore() != 71 {
		t.Fatalf("Auction score for a new board should be 71 but is %d", b.AuctionScore())
	}
}

func TestBoardSideDeckIsEmptyAtCreation(t *testing.T) {
	if b := New(); len(b.side) != 0 {
		t.Fatal("Size of side deck should be 0")
	}
}

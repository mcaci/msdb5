package board

import (
	"testing"
)

func TestActionCreationAndJoinUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || board.Players()[0].Name() != "A" {
		t.Fatal("Join action was not properly performed")
	}
}

func TestActionCreationAndAuctionUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	b.Action("Auction#102", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || board.AuctionScore() != 102 {
		t.Fatal("Auction action was not properly performed")
	}
}

func TestActionCreationAndPCompanionUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	b.Action("Companion#3#Cup", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || *board.NominatedCard() != 13 {
		t.Fatal("Companion action was not properly performed")
	}
}
func TestActionCreationAndPlayCardUsage(t *testing.T) {
	b := NewAction()
	b.Action("Join#A", "100.1.1.1")
	b.Action("Card#6#Cudgel", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || !board.PlayedCards().Has(36) {
		t.Fatal("Card action was not properly performed")
	}
}

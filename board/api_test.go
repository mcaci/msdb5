package board

import (
	"testing"
)

func TestAPICreationAndJoinUsage(t *testing.T) {
	b := NewAPI()
	b.Action("Join#A", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || board.Players()[0].Name() != "A" {
		t.Fatal("Join action was not properly performed")
	}
}

func TestAPICreationAndAuctionUsage(t *testing.T) {
	b := NewAPI()
	b.Action("Join#A", "100.1.1.1")
	b.Action("Auction#102", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || board.AuctionScore() != 102 {
		t.Fatal("Auction action was not properly performed")
	}
}

func TestAPICreationAndPlayUsage(t *testing.T) {
	b := NewAPI()
	b.Action("Join#A", "100.1.1.1")
	b.Action("Play#6#Cudgel", "100.1.1.1")
	board, ok := b.(*Board)
	if !ok || *board.NominatedCard() != 36 {
		t.Fatal("Play action was not properly performed")
	}
}

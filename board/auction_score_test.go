package board

import (
	"testing"
)

func TestBoardAuctionScoreAtCreationIs0(t *testing.T) {
	if b := New(); b.AuctionScore() != 0 {
		t.Fatalf("Auction score for a new board should be 0 but is %d", b.AuctionScore())
	}
}

func TestBoardAuctionScoreCanBeSet(t *testing.T) {
	b := New()
	b.SetAuctionScore(80)
	if b.AuctionScore() != 80 {
		t.Fatalf("Auction score should be set at 80 but is %d", b.AuctionScore())
	}
}

func TestRaiseAuctionScoreFirstAssignment(t *testing.T) {
	b := New()
	if b.RaiseAuction("61"); b.AuctionScore() != 61 {
		t.Fatalf("Auction score should be set at 61 but is %d", b.AuctionScore())
	}
}

func TestRaiseAuctionScoreFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	b := New()
	if b.RaiseAuction("1"); b.AuctionScore() != 61 {
		t.Fatalf("Auction score should be set at 61 but is %d", b.AuctionScore())
	}
}
func TestInvalidRaiseAuctionScoreFirstAssignmentShouldBeAlways61(t *testing.T) {
	b := New()
	if b.RaiseAuction("ciao"); b.AuctionScore() != 61 {
		t.Fatalf("Auction score should be set at 61 but is %d", b.AuctionScore())
	}
}

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

func TestRaiseAuctionScoreSecondAssignmentShouldBeSuperiorThanFirstOne(t *testing.T) {
	b := New()
	b.RaiseAuction("65")
	if b.RaiseAuction("80"); b.AuctionScore() != 80 {
		t.Fatalf("Auction score should be set at 80 but is %d", b.AuctionScore())
	}
}

func TestRaiseAuctionScoreSecondAssignmentShouldBeSuperiorThanFirstOneElseDrop(t *testing.T) {
	b := New()
	b.RaiseAuction("65")
	if b.RaiseAuction("63"); b.AuctionScore() != 65 {
		t.Fatalf("Auction score should be set at 65 but is %d", b.AuctionScore())
	}
}

func TestInvalidSecondRaiseAuctionScoreAlwaysDrops(t *testing.T) {
	b := New()
	b.RaiseAuction("90")
	if b.RaiseAuction("ciao"); b.AuctionScore() != 90 {
		t.Fatalf("Auction score should be set at 90 but is %d", b.AuctionScore())
	}
}

func TestRaiseAuctionCannotPass120Score(t *testing.T) {
	b := New()
	b.RaiseAuction("99")
	if b.RaiseAuction("125"); b.AuctionScore() != 120 {
		t.Fatalf("Auction score should be set at 120 but is %d", b.AuctionScore())
	}
}

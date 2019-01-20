package board

import (
	"testing"
)

func testPlayerScore(t *testing.T, expectedScore, actualScore uint8) {
	if expectedScore != actualScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}

func TestRaiseAuctionScoreWithHostFirstAssignment(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction2("61", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 61)
}

func TestRaiseAuctionScoreWithHostFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction2("1", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 61)
}

func TestInvalidRaiseAuctionScoreWithHostFirstAssignmentShouldBeAlways61(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction2("ciao", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 61)
}

func TestRaiseAuctionTo65WithHost(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction2("65", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 65)

}

// func Test2Players(t *testing.T) {
// 	b := New()
// 	b.RaiseAuction2("65", "")
// 	if b.RaiseAuction2("80", ""); b.AuctionScore() != 80 {
// 		t.Fatalf("Auction score should be set at 80 but is %d", b.AuctionScore())
// 	}
// }

// func TestRaiseAuctionScoreSecondAssignmentShouldBeSuperiorThanFirstOneElseDrop(t *testing.T) {
// 	b := New()
// 	b.RaiseAuction2("65", "")
// 	if b.RaiseAuction2("63", ""); b.AuctionScore() != 65 {
// 		t.Fatalf("Auction score should be set at 65 but is %d", b.AuctionScore())
// 	}
// }

// func TestInvalidSecondRaiseAuctionScoreAlwaysDrops(t *testing.T) {
// 	b := New()
// 	b.RaiseAuction2("90", "")
// 	if b.RaiseAuction2("ciao", ""); b.AuctionScore() != 90 {
// 		t.Fatalf("Auction score should be set at 90 but is %d", b.AuctionScore())
// 	}
// }

// func TestRaiseAuctionCannotPass120Score(t *testing.T) {
// 	b := New()
// 	b.RaiseAuction2("99", "")
// 	if b.RaiseAuction2("125", ""); b.AuctionScore() != 120 {
// 		t.Fatalf("Auction score should be set at 120 but is %d", b.AuctionScore())
// 	}
// }

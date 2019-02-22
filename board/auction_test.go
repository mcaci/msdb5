package board

import (
	"testing"
)

func testPlayerScore(t *testing.T, actualScore, expectedScore uint8) {
	if actualScore != expectedScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}

func TestRaiseAuctionScoreWithHostFirstAssignment(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction("61", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 61)
}

func TestRaiseAuctionScoreWithHostFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction("1", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 61)
}

func TestInvalidRaiseAuctionScoreWithHostFirstAssignmentShouldBeAlways61(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction("ciao", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 61)
}

func TestRaiseAuctionTo65WithHost(t *testing.T) {
	b := New()
	b.Join("name", "100.1.1.1")
	b.RaiseAuction("65", "100.1.1.1")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 65)
}

func Test2PlayersRaisingAuctionWithHost(t *testing.T) {
	b := New()
	b.Join("A", "100.1.1.1")
	b.Join("B", "100.1.1.2")
	b.RaiseAuction("65", "100.1.1.1")
	b.RaiseAuction("80", "100.1.1.2")
	testPlayerScore(t, b.Players()[0].AuctionScore(), 65)
}

func Test2PlayersRaisingAuctionWithHostSecondPlayer(t *testing.T) {
	b := New()
	b.Join("A", "100.1.1.1")
	b.Join("B", "100.1.1.2")
	b.RaiseAuction("65", "100.1.1.1")
	b.RaiseAuction("80", "100.1.1.2")
	testPlayerScore(t, b.Players()[1].AuctionScore(), 80)
}

func Test2PlayersRaisingAuctionWithHostSecondPlayerDropsWithLowerScore(t *testing.T) {
	b := New()
	b.Join("A", "100.1.1.1")
	b.Join("B", "100.1.1.2")
	b.RaiseAuction("65", "100.1.1.1")
	b.RaiseAuction("61", "100.1.1.2")
	testPlayerScore(t, b.Players()[1].AuctionScore(), 0)
}

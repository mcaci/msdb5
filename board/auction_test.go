package board

import (
	"testing"
)

func MockBoard() *Board {
	testBoard := New()
	testBoard.Join("A", "100.1.1.1")
	testBoard.Join("B", "100.1.1.2")
	return testBoard
}

func TestRaiseAuctionScoreFirstAssignment(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("61", "100.1.1.1")
	testPlayerScore(t, testBoard.Players()[0].AuctionScore(), 61)
}

func TestRaiseAuctionScoreFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("1", "100.1.1.1")
	testPlayerScore(t, testBoard.Players()[0].AuctionScore(), 61)
}

func TestInvalidRaiseAuctionScoreFirstAssignmentShouldBeAlways61(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("ciao", "100.1.1.1")
	testPlayerScore(t, testBoard.Players()[0].AuctionScore(), 61)
}

func TestRaiseAuctionTo65(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("65", "100.1.1.1")
	testPlayerScore(t, testBoard.Players()[0].AuctionScore(), 65)
}

func Test2PlayersRaisingAuction(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("65", "100.1.1.1")
	testBoard.RaiseAuction("80", "100.1.1.2")
	testPlayerScore(t, testBoard.Players()[0].AuctionScore(), 65)
}

func Test2PlayersRaisingAuctionSecondPlayer(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("65", "100.1.1.1")
	testBoard.RaiseAuction("80", "100.1.1.2")
	testPlayerScore(t, testBoard.Players()[1].AuctionScore(), 80)
}

func Test2PlayersRaisingAuctionSecondPlayerDropsWithLowerScore(t *testing.T) {
	testBoard := MockBoard()
	testBoard.RaiseAuction("65", "100.1.1.1")
	testBoard.RaiseAuction("61", "100.1.1.2")
	testPlayerScore(t, testBoard.Players()[1].AuctionScore(), 0)
}

func testPlayerScore(t *testing.T, actualScore, expectedScore uint8) {
	if actualScore != expectedScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}

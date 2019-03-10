package orchestrator

import (
	"testing"
)

func mockGameAuctionTest() *Game {
	gameTest := NewGame()
	gameTest.players[0].Join("A", "127.0.0.11")
	gameTest.players[1].Join("B", "127.0.0.12")
	gameTest.players[2].Join("C", "127.0.0.13")
	gameTest.players[3].Join("D", "127.0.0.14")
	gameTest.players[4].Join("E", "127.0.0.15")
	gameTest.phase = scoreAuction
	return gameTest
}

func TestPlayerCannotRaiseAuctionIfPhaseIsNotAuction(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.phase = joining
	_, _, err := gameTest.raiseAuction("Auction#102", "127.0.0.11")
	if err == nil {
		t.Fatal("Auction action not expected at beginning of game")
	}
}

func TestFirstPlayerCanRaiseAuction(t *testing.T) {
	gameTest := mockGameAuctionTest()
	_, _, err := gameTest.raiseAuction("Auction#102", "127.0.0.11")
	if err != nil {
		t.Fatal("Expecting first player to raise auction with success")
	}
}

func TestSecondPlayerCannotRaiseAuctionIfNotDoneByFirstPlayer(t *testing.T) {
	gameTest := mockGameAuctionTest()
	_, _, err := gameTest.raiseAuction("Auction#97", "127.0.0.12")
	if err == nil {
		t.Fatal("Expecting error for second player not being able to act before first player has raised the auction")
	}
}

func TestSecondPlayerCanRaiseAuctionAfterFirstPlayer(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.raiseAuction("Auction#78", "127.0.0.11")
	_, _, err := gameTest.raiseAuction("Auction#81", "127.0.0.12")
	if err != nil {
		t.Fatal("Second player should be able to act after first player has raised the auction")
	}
}

func TestSecondPlayerCanFoldAuctionAfterFirstPlayer(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.raiseAuction("Auction#98", "127.0.0.11")
	_, _, err := gameTest.raiseAuction("Auction#ciao", "127.0.0.12")
	if err != nil {
		t.Fatal("Second player should be able to act after first player has raised the auction")
	}
}

func TestSkipPlayerThatHasFolded(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.players[1].Fold()
	gameTest.raiseAuction("Auction#80", "127.0.0.11")
	_, _, err := gameTest.raiseAuction("Auction#85", "127.0.0.13")
	if err != nil {
		t.Fatal("Folded player, player 2, was not skipped")
	}
}

func TestGoToNominateWhenAuctionEnds(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.players[1].Fold()
	gameTest.players[2].Fold()
	gameTest.players[4].Fold()
	gameTest.raiseAuction("Auction#80", "127.0.0.11")
	gameTest.raiseAuction("Auction#85", "127.0.0.14")
	gameTest.raiseAuction("Auction#ciao", "127.0.0.11")
	if gameTest.phase != companionChoice {
		t.Fatal("Auction round is over but game did not step to the companion choice phase")
	}
}

func TestAuctionWinnerSelectionWhenAuctionEnds(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.players[1].Fold()
	gameTest.players[2].Fold()
	gameTest.players[4].Fold()
	gameTest.raiseAuction("Auction#80", "127.0.0.11")
	gameTest.raiseAuction("Auction#85", "127.0.0.14")
	gameTest.raiseAuction("Auction#ciao", "127.0.0.11")
	if gameTest.playerInTurn != 3 {
		t.Fatalf("D should be the auction winner but was %d", gameTest.playerInTurn)
	}
}

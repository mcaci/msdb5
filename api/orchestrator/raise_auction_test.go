package orchestrator

import (
	"testing"
)

func mockGameAuctionTest() *Game {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.Join("C", "100.1.1.3")
	gameTest.Join("D", "100.1.1.4")
	gameTest.Join("E", "100.1.1.5")
	gameTest.phase = scoreAuction
	return gameTest
}

func TestPlayerCannotRaiseAuctionIfPhaseIsNotAuction(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.phase = joining
	err := gameTest.RaiseAuction("102", "100.1.1.1")
	if err == nil {
		t.Fatal("Auction action not expected at beginning of game")
	}
}

func TestFirstPlayerCanRaiseAuction(t *testing.T) {
	gameTest := mockGameAuctionTest()
	err := gameTest.RaiseAuction("102", "100.1.1.1")
	if err != nil {
		t.Fatal("Expecting first player to raise auction with success")
	}
}

func TestSecondPlayerCannotRaiseAuctionIfNotDoneByFirstPlayer(t *testing.T) {
	gameTest := mockGameAuctionTest()
	err := gameTest.RaiseAuction("97", "100.1.1.2")
	if err == nil {
		t.Fatal("Expecting error for second player not being able to act before first player has raised the auction")
	}
}

func TestSecondPlayerCanRaiseAuctionAfterFirstPlayer(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.RaiseAuction("78", "100.1.1.1")
	err := gameTest.RaiseAuction("81", "100.1.1.2")
	if err != nil {
		t.Fatal("Second player should be able to act after first player has raised the auction")
	}
}

func TestSecondPlayerCanFoldAuctionAfterFirstPlayer(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.RaiseAuction("98", "100.1.1.1")
	err := gameTest.RaiseAuction("ciao", "100.1.1.2")
	if err != nil {
		t.Fatal("Second player should be able to act after first player has raised the auction")
	}
}

func TestSkipPlayerThatHasFolded(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.players[1].Fold()
	gameTest.RaiseAuction("80", "100.1.1.1")
	err := gameTest.RaiseAuction("85", "100.1.1.3")
	if err != nil {
		t.Fatal("Folded player, player 2, was not skipped")
	}
}

func TestGoToNominateWhenAuctionEnds(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.players[1].Fold()
	gameTest.players[2].Fold()
	gameTest.players[4].Fold()
	gameTest.RaiseAuction("80", "100.1.1.1")
	gameTest.RaiseAuction("85", "100.1.1.4")
	gameTest.RaiseAuction("ciao", "100.1.1.1")
	if gameTest.phase != companionChoice {
		t.Fatal("Auction round is over but game did not step to the companion choice phase")
	}
}

func TestAuctionWinnerSelectionWhenAuctionEnds(t *testing.T) {
	gameTest := mockGameAuctionTest()
	gameTest.players[1].Fold()
	gameTest.players[2].Fold()
	gameTest.players[4].Fold()
	gameTest.RaiseAuction("80", "100.1.1.1")
	gameTest.RaiseAuction("85", "100.1.1.4")
	gameTest.RaiseAuction("ciao", "100.1.1.1")
	if gameTest.playerInTurn != 3 {
		t.Fatalf("D should be the auction winner but was %d", gameTest.playerInTurn)
	}
}

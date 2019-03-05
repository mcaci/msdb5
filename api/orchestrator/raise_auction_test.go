package orchestrator

import "testing"

func TestPlayerCannotRaiseAuctionIfPhaseIsNotAuction(t *testing.T) {
	gameTest := NewGame()
	err := gameTest.RaiseAuction("102", "100.1.1.1")
	if err == nil {
		t.Fatal("Auction action not expected at beginning of game")
	}
}

func TestPlayerCannotRaiseAuctionIfPhaseIsNotAuctionEvenAfterFirstJoin(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	err := gameTest.RaiseAuction("102", "100.1.1.1")
	if err == nil {
		t.Fatal("Auction action not expected at beginning of game")
	}
}

func TestFirstPlayerCanRaiseAuction(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.phase = scoreAuction
	err := gameTest.RaiseAuction("102", "100.1.1.1")
	if err != nil {
		t.Fatal("Expecting first player to raise auction with success")
	}
}

func TestSecondPlayerCannotRaiseAuctionIfNotDoneByFirstPlayer(t *testing.T) {
	gameTest := NewGame()
	gameTest.Join("A", "100.1.1.1")
	gameTest.Join("B", "100.1.1.2")
	gameTest.phase = scoreAuction
	err := gameTest.RaiseAuction("97", "100.1.1.2")
	if err == nil {
		t.Fatal("Expecting error for second player not being able to act before first player has raised the auction")
	}
}

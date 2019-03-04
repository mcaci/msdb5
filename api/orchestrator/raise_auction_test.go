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

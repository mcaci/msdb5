package board

import "testing"

func TestPlayerPlaysOneCardAndCardIsRemovedFromHand(t *testing.T) {
	// TODO: fix this test
	b := New()
	b.Join("A", "100.0.0.1")
	b.Play("1", "Coin", "100.0.0.1")
	if !b.PlayedCards().Has(1) {
		t.Fatal("Board should have One of Coin as played card")
	}
}

func TestPlayerPlaysOneCardAndCardIsOnTheBoard(t *testing.T) {
	b := New()
	b.Join("A", "100.0.0.1")
	b.Play("1", "Coin", "100.0.0.1")
	if !b.PlayedCards().Has(1) {
		t.Fatal("Board should have One of Coin as played card")
	}
}

package board

import "testing"

func TestPlayerPlaysOneCardAndCardIsRemovedFromHand(t *testing.T) {
	b := New()
	card := playCard(b)
	if b.Players()[0].Has(card) {
		t.Fatalf("Hand should not have %v", card)
	}
}

func TestPlayerPlaysOneCardAndCardIsOnTheBoard(t *testing.T) {
	b := New()
	card := playCard(b)
	if !b.PlayedCards().Has(card) {
		t.Fatalf("Played cards should have %v", card)
	}
}

package board

import (
	"testing"
)

func TestBoardHasASetOfPlayedCards(t *testing.T) {
	if b := New(); b.PlayedCards() == nil {
		t.Fatal("The board has no set of played cards")
	}
}

func TestBoardsEmptySetOfPlayedCardsContainsNoCards(t *testing.T) {
	if b := New(); b.PlayedCards().Has(1) {
		t.Fatal("The deck should be empty at this point")
	}
}

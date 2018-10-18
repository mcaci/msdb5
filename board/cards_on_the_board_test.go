package board

import (
	"testing"
)

func TestBoardHasASetOfPlayedCards(t *testing.T) {
	b := New()
	if b.PlayedCards() == nil {
		t.Fatal("The board has no set of played cards")
	}

}

func TestBoardsEmptySetOfPlayedCardsContainsNoCards(t *testing.T) {
	b := New()
	if b.PlayedCards().HasID(1) {
		t.Fatal("The deck should be empty at this point")
	}
}

func TestBoardsSetOfPlayedCardsWithOneCardContainsIt(t *testing.T) {
	b := New()
	b.PlayedCards().AddID(1)
	if !b.PlayedCards().HasID(1) {
		t.Fatal("The deck should contain one card")
	}
}

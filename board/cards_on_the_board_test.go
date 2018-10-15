package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestBoardHasASetOfPlayedCards(t *testing.T) {
	b := New()
	if b.PlayedCards() == nil {
		t.Fatal("The board has no set of played cards")
	}

}

func TestBoardsEmptySetOfPlayedCardsContainsNoCards(t *testing.T) {
	b := New()
	c, _ := card.ByID(1)
	if b.PlayedCards().Has(c) {
		t.Fatal("The deck should be empty at this point")
	}
}

func TestBoardsSetOfPlayedCardsWithOneCardContainsIt(t *testing.T) {
	b := New()
	c, _ := card.ByID(1)
	b.playedCards = append(b.playedCards, c)
	if !b.PlayedCards().Has(c) {
		t.Fatal("The deck should contain one card")
	}
}

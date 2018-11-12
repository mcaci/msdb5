package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
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

func TestBoardsSetOfPlayedCardsWithOneCardContainsIt(t *testing.T) {
	b := New()
	b.PlayedCards().Add(1)
	if !b.PlayedCards().Has(1) {
		t.Fatal("The deck should contain one card")
	}
}

func TestPlayerPlaysOneCardAndCardIsOnTheBoard(t *testing.T) {
	b := New()
	h := b.Players()[0].Hand()
	card := (*h)[0]
	removeCardFromH(card, h)
	b.PlayedCards().Add(card)
	if h.Has(card) {
		t.Fatalf("Hand should not have %v", card)
	}
}

func removeCardFromH(c card.ID, h *card.Cards) {
	index := 0
	for i, card := range *h {
		if card == c {
			index = i
			break
		}
	}
	*h = append((*h)[:index], (*h)[index+1])
}

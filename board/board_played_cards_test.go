package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/card/set"
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

func playCard(b *Board) card.ID {
	h := b.Players()[0].Hand()
	card := (*h)[0]
	removeCardFromH(card, h)
	b.PlayedCards().Add(card)
	return card
}

func removeCardFromH(c card.ID, h *set.Cards) {
	index := 0
	for i, card := range *h {
		if card == c {
			index = i
			break
		}
	}
	*h = append((*h)[:index], (*h)[index+1])
}

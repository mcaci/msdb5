package play

import (
	"testing"

	"github.com/mcaci/msdb5/dom/deck"
)

func TestExchangeNotFailing(t *testing.T) {
	pCards := deck.Cards{1}
	side := deck.Cards{2}
	err := CardAction(fakeInput(1), &pCards, &side, func(cards, to *deck.Cards, index, toIndex int) {
		(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCardsAreExchanged(t *testing.T) {
	pCards := deck.Cards{1}
	side := deck.Cards{2}
	CardAction(fakeInput(1), &pCards, &side, func(cards, to *deck.Cards, index, toIndex int) {
		(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
	})
	if side[0] != 1 {
		t.Fatalf("Expecting 1 to be present in side deck but was %v", side[0])
	}
}

func TestExchangeFailing(t *testing.T) {
	pCards := deck.Cards{1}
	side := deck.Cards{2}
	err := CardAction(fakeInput(3), &pCards, &side, func(cards, to *deck.Cards, index, toIndex int) {
		(*cards)[index], (*to)[toIndex] = (*to)[index], (*cards)[toIndex]
	})
	if err == nil {
		t.Fatal("An error was supposed to be returned")
	}
}

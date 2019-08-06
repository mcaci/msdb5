package game

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
)

func TestPostExchangeCardsResult(t *testing.T) {
	cards := &deck.Cards{1, 2, 3}
	to := &deck.Cards{4, 5, 6, 7, 8}
	index := 2
	postExchange(cards, to, index)
	if (*cards)[index] != 4 {
		t.Fatalf("Expecting %s, found %s", card.ID(4), (*cards)[index])
	}
}

func TestPostExchangeToResult(t *testing.T) {
	cards := &deck.Cards{1, 2, 3}
	to := &deck.Cards{4, 5, 6, 7, 8}
	index := 2
	postExchange(cards, to, index)
	if (*to)[len(*to)-1] != 3 {
		t.Fatalf("Expecting %s, found %s", card.ID(3), (*to)[len(*to)-1])
	}
}

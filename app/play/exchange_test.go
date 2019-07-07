package play

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

func TestExchangeNotFailing(t *testing.T) {
	pCards := deck.Cards{1}
	side := deck.Cards{2}
	err := Exchange(1, &pCards, &side)
	if err != nil {
		t.Error(err)
	}
}

func TestCardsAreExchanged(t *testing.T) {
	pCards := deck.Cards{1}
	side := deck.Cards{2}
	Exchange(1, &pCards, &side)
	if side[0] != 1 {
		t.Fatalf("Expecting 1 to be present in side deck but was %v", side[0])
	}
}

func TestExchangeFailing(t *testing.T) {
	pCards := deck.Cards{1}
	side := deck.Cards{2}
	err := Exchange(3, &pCards, &side)
	if err == nil {
		t.Fatal("An error was supposed to be returned")
	}
}

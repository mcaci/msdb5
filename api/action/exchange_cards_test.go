package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/api/game"
)

func TestExchangeCardsNextPlayerOf2is2(t *testing.T) {
	if testObject := NewExchangeCards("", "", nil); testObject.NextPlayer(2) != 2 {
		t.Fatalf("Next player should be 2")
	}
}

func TestExchangeCardsNextPlayerOf4is4(t *testing.T) {
	if testObject := NewExchangeCards("", "", nil); testObject.NextPlayer(4) != 4 {
		t.Fatalf("Next player should be 1")
	}
}

func TestExchangeCardsNextPhaseWhenInputIs0(t *testing.T) {
	if testObject := NewExchangeCards("#0", "", nil); game.ChosingCompanion != testObject.NextPhase() {
		t.Fatalf("Should change phase when 0 is in the request")
	}
}

func TestExchangeCardsNextPhaseWhenInputIsNot0(t *testing.T) {
	if testObject := NewExchangeCards("#1", "", nil); game.ExchangingCards != testObject.NextPhase() {
		t.Fatalf("Should not change phase when 1 is in the request")
	}
}
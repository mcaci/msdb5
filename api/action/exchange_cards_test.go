package action

import (
	"github.com/nikiforosFreespirit/msdb5/api/game"
	"testing"
)

func TestExchangeCardsPhase(t *testing.T) {
	if testObject := NewExchangeCards("", ""); testObject.Phase() != game.ExchangingCards {
		t.Fatalf("Unexpected phase")
	}
}
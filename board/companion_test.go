package board

import (
	"testing"
)

func TestCardNominatedWithOriginInfo(t *testing.T) {
	b := New()
	b.Nominate("1", "Coin", "100.1.1.1")
	if *b.NominatedCard() != 1 {
		t.Fatal("Nominated card is not 1 of Coin")
	}
}

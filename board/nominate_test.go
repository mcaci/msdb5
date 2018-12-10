package board

import (
	"testing"
)

func TestNominateId1WillNominateAceOfCoin(t *testing.T) {
	board := New()
	actualCard := board.AskNominatedCard()
	expectedCard := board.NominatedCard()
	if *expectedCard != actualCard {
		t.Fatalf("Card nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

func TestCardNominatedIsStoredOnBoard(t *testing.T) {
	board := New()
	actualCard, err := board.Nominate("1", "Coin")
	expectedCard := *board.NominatedCard()
	if err != nil || expectedCard != actualCard {
		t.Fatalf("Card nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

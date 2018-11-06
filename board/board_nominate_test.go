package board

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestNominateId0WillComplain(t *testing.T) {
	board := New()
	if _, err := board.Nominate("0", "Coin"); err == nil {
		t.Fatal("Error should be present")
	}
}

func TestNominateId1WillNominateAceOfCoin(t *testing.T) {
	board := New()
	actualCard, err := board.Nominate("1", "Coin")
	expectedCard, _ := card.ByName("1", "Coin")
	if err != nil || expectedCard != actualCard {
		t.Fatalf("Data nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

func TestCardNominatedIsStoredOnBoard(t *testing.T) {
	board := New()
	actualCard, err := board.Nominate("1", "Coin")
	expectedCard := *board.NominatedCard()
	if err != nil || expectedCard != actualCard {
		t.Fatalf("Data nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

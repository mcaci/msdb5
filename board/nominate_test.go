package score

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestNominateId0WillComplain(t *testing.T) {
	if _, err := Nominate("0", "Coin"); err == nil {
		t.Fatal("Error should be present")
	}
}

func TestNominateId1WillNominateAceOfCoin(t *testing.T) {
	actualCard, err := Nominate("1", "Coin")
	expectedCard, _ := card.ByName("1", "Coin")
	if err != nil || expectedCard != actualCard {
		t.Fatalf("Card nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

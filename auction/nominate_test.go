package score

import (
	"testing"
	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestNominateId0WillComplain(t *testing.T) {
	_, err := Nominate(0)
	if err == nil {
		t.Fatal("Error should be present")
	}
}

func TestNominateId1WillNominateAceOfCoin(t *testing.T) {
	actualCard, err := Nominate(1)
	expectedCard, _ := card.ByID(1)
	if err != nil || *expectedCard != *actualCard {
		t.Fatalf("Card nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

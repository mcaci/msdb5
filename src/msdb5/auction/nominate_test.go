package score

import (
	"testing"
	"msdb5/card"
)

func Nominate(id uint8) *card.Card {
	return nil
}

func TestNominateId1WillNominateAceOfCoin(t *testing.T) {
	actualCard := Nominate(1)
	expectedCard, _ := card.ByID(1)
	if expectedCard != actualCard {
		t.Fatalf("Card nominated should be %v but %d was computed", expectedCard, actualCard)
	}
}

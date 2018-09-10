package score

import (
	"testing"
	"msdb5/card"
)

func Nominate(id uint8) *card.Card {
	card, _ := card.ByID(id)
	return card
}

func TestNominateId1WillNominateAceOfCoin(t *testing.T) {
	actualCard := Nominate(1)
	expectedCard, _ := card.ByID(1)
	if *expectedCard != *actualCard {
		t.Fatalf("Card nominated should be %v but %v was computed", expectedCard, actualCard)
	}
}

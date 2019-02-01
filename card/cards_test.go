package card

import "testing"

func TestCreateSet(t *testing.T) {
	cards := Cards{15}
	if !cards.Has(15) {
		t.Fatalf("There should be the 5 of Cup card in the set")
	}
}

func TestRemoveCardFromSet(t *testing.T) {
	cards := Cards{15}
	cards.Remove(0)
	if cards.Has(15) {
		t.Fatalf("Deck should be empty")
	}
}

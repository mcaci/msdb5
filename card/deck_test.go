package card

import "testing"

func TestCreate(t *testing.T) {
	d := Deck()
	count := 0
	for !d.IsEmpty() {
		count++
		d.Supply()
	}
	if count != DeckSize {
		t.Fatalf("There should be 40 card in the deck")
	}
}

func TestRemovingTwoCardsShouldGiveDifferentCards(t *testing.T) {
	d := Deck()
	a := d.Supply()
	b := d.Supply()

	if a == b {
		t.Fatalf("Drawn cards should be different but they are %v and %v", a, b)
	}
}

package deck

import "testing"

func TestCreateDeck(t *testing.T) {
	d := New()
	if len(d) != DeckSize {
		t.Fatalf("There should be 40 card in the deck")
	}
}

func TestRemovingTwoCardsShouldGiveDifferentCards(t *testing.T) {
	d := New()
	a := d.Supply()
	b := d.Supply()

	if a == b {
		t.Fatalf("Drawn cards should be different but they are %v and %v", a, b)
	}
}

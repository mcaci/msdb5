package deck

import "testing"

func TestCreate(t *testing.T) {
	d := New()
	count := 0
	for !d.IsEmpty() {
		count++
		d.Supply()
	}
	if count != Size {
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

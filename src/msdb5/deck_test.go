package msdb5

import "testing"

func createDeck() Deck {
	var d = &ConcreteDeck{}
	d.Create()
	return d
}

func TestCreate(t *testing.T) {
	var d = createDeck()
	sum := 0
	for _, v := range d.(*ConcreteDeck).cards {
		sum += v
	}
	if sum != 780 {
		t.Fatalf("Cards id sum should be the sum of all numbers up to 39")
	}
}

func TestRemoveTop(t *testing.T) {
	var d = createDeck()
	a := d.RemoveTop()
	b := d.RemoveTop()

	if a == b {
		t.Fatalf("Drawn cards should be different but they are %v and %v", a, b)
	}
}

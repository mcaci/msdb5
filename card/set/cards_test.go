package set

import "testing"

func TestCreateSet(t *testing.T) {
	cards := Cards{15}
	if !cards.Has(15) {
		t.Fatalf("There should be the 5 of Cup card in the set")
	}
}

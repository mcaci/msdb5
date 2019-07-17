package deck

import (
	"testing"
)

func TestCreateSet(t *testing.T) {
	set := Cards{15}
	if index := set.Find(15); index == -1 {
		t.Fatalf("There should be the 5 of Cup card in the set")
	}
}

func TestAddCardToSet(t *testing.T) {
	set := Cards{}
	set.Add(33)
	if index := set.Find(33); index == -1 {
		t.Fatal("There should be the 3 of Cudgel card in the set")
	}
}

func TestClearRemovesAllCards(t *testing.T) {
	set := Cards{2}
	set.Clear()
	if len(set) != 0 {
		t.Fatalf("Cards were not removed from set")
	}
}

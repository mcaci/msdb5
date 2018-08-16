package card

import "testing"

func TestId0IsAceOfCoin(t *testing.T) {
	card := ByID(1)
	if card.number != 1 {
		t.Fatalf("Card %v is not an ace", card)
	}
}

// func TestId39IsKingOfSomething(t *testing.T) {
// 	card := ByID(39)
// 	if card.number != 9 {
// 		t.Fatalf("Card %v is not a nine", card)
// 	}
// }

// func ByID(id int) *Card {
// 	a := uint8(id % 10)
// 	b := Seed(id / 10)
// 	return &Card{number: a, seed: b}
// }

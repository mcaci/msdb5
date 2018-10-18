package card

import (
	"testing"
)

func TestId1IsAce(t *testing.T) {
	verifyCorrectNumber(t, 1, 1)
}

func TestId12IsTwo(t *testing.T) {
	verifyCorrectNumber(t, 12, 2)
}

func TestId25IsFive(t *testing.T) {
	verifyCorrectNumber(t, 25, 5)
}

func TestId40IsKing(t *testing.T) {
	verifyCorrectNumber(t, 40, 10)
}

func TestId0IsInvalid(t *testing.T) {
	verifyInvalidID(t, 0)
}
func TestId41IsInvalid(t *testing.T) {
	verifyInvalidID(t, 41)
}

func TestId1IsCoin(t *testing.T) {
	verifyCorrectSeed(t, 1, Coin)
}

func TestId12IsCup(t *testing.T) {
	verifyCorrectSeed(t, 12, Cup)
}

func TestId25IsSword(t *testing.T) {
	verifyCorrectSeed(t, 25, Sword)
}

func TestId40IsCudgel(t *testing.T) {
	verifyCorrectSeed(t, 40, Cudgel)
}

func TestFromIdToCardToId(t *testing.T) {
	card, _ := ByID(1)
	if 1 != card.ID() {
		t.Fatalf("Card ids are not the same")
	}
}

func verifyCorrectSeed(t *testing.T, id ID, seed Seed) {
	card, _ := ByID(id)
	if card.seed != seed {
		t.Fatalf("Card %v's seed is not %s", card, seed)
	}
}

func verifyInvalidID(t *testing.T, id ID) {
	_, err := ByID(id)
	if err == nil {
		t.Fatalf("%d is not valid id", id)
	}
}

func verifyCorrectNumber(t *testing.T, id ID, number uint8) {
	card, _ := ByID(id)
	if card.number != number {
		t.Fatalf("Card %v's number is not %d", card, number)
	}
}

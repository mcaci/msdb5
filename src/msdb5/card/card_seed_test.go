package card

import "testing"

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

func verifyCorrectSeed(t *testing.T, id int, seed Seed) {
	card, _ := ByID(id)
	if card.seed != seed {
		t.Fatalf("Card %v's seed is not %s", card, seed)
	}
}

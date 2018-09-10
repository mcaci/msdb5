package card

import "testing"

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

func verifyCorrectNumber(t *testing.T, id, number uint8) {
	card, _ := ByID(id)
	if card.number != uint8(number) {
		t.Fatalf("Card %v's number is not %d", card, number)
	}
}

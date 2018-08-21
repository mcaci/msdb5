package card

import "testing"

func TestId1IsAce(t *testing.T) {
	card, _ := ByID(1)
	if card.number != 1 {
		t.Fatalf("Card %v's number is not ace", card)
	}
}

func TestId1IsCoin(t *testing.T) {
	card, _ := ByID(1)
	if card.seed != Coin {
		t.Fatalf("Card %v's seed is not coin", card)
	}
}

func TestId12IsTwo(t *testing.T) {
	card, _ := ByID(12)
	if card.number != 2 {
		t.Fatalf("Card %v's number is not two", card)
	}
}

func TestId12IsCup(t *testing.T) {
	card, _ := ByID(12)
	if card.seed != Cup {
		t.Fatalf("Card %v's seed is not cup", card)
	}
}

func TestId25IsFive(t *testing.T) {
	card, _ := ByID(25)
	if card.number != 5 {
		t.Fatalf("Card %v's number is not five", card)
	}
}

func TestId25IsSword(t *testing.T) {
	card, _ := ByID(25)
	if card.seed != Sword {
		t.Fatalf("Card %v's seed is not sword", card)
	}
}

func TestId40IsKing(t *testing.T) {
	card, _ := ByID(40)
	if card.number != 10 {
		t.Fatalf("Card %v's number is not king", card)
	}
}

func TestId40IsCudgel(t *testing.T) {
	card, _ := ByID(40)
	if card.seed != Cudgel {
		t.Fatalf("Card %v's seed is not cudgel", card)
	}
}

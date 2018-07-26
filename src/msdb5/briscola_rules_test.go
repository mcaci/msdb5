package msdb5

import "testing"

func TestInClassicNumericalComparisonWithSameSeedHigherNumberWins(t *testing.T) {
	a := Card{number: 2, seed: Coin}
	b := Card{number: 4, seed: Coin}
	c := higherCardBetween(&a, &b)
	if c != &b {
		t.Fatalf("Expected %v to be higher than %v. %v was the output", b, a, *c)
	}
}

func TestInComparisonWithSameSeedThat3isHigherThan10(t *testing.T) {
	a := Card{number: 10, seed: Coin}
	b := Card{number: 3, seed: Coin}
	c := higherCardBetween(&a, &b)
	if c != &b {
		t.Fatalf("Expected %v to be higher than %v. %v was the output", b, a, *c)
	}
}

func TestInComparisonWithSameSeedThat1isHigherThan8(t *testing.T) {
	a := Card{number: 1, seed: Coin}
	b := Card{number: 9, seed: Coin}
	c := higherCardBetween(&a, &b)
	if c != &a {
		t.Fatalf("Expected %v to be higher than %v. %v was the output", a, b, *c)
	}
}


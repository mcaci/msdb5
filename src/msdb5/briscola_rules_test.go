package msdb5

import "testing"

var firstCardIsHigher = func(i int) bool { return i > 0 }
var firstCardIsLower = func(i int) bool { return i < 0 }

func TestInClassicNumericalComparisonWithSameSeedHigherNumberWins(t *testing.T) {
	a := Card{number: 2, seed: Coin}
	b := Card{number: 4, seed: Coin}
	verify(t, &a, &b, firstCardIsLower)
}

func TestInComparisonWithSameSeedThat3isHigherThan10(t *testing.T) {
	a := Card{number: 10, seed: Coin}
	b := Card{number: 3, seed: Coin}
	verify(t, &a, &b, firstCardIsLower)
}

func TestInComparisonWithSameSeedThat1isHigherThan8(t *testing.T) {
	a := Card{number: 1, seed: Coin}
	b := Card{number: 9, seed: Coin}
	verify(t, &a, &b, firstCardIsHigher)
}

func verify(t *testing.T, a, b *Card, isComparisonBetweenCardsCorrect func(int) bool) {
	c := a.Compare(*b)
	if !isComparisonBetweenCardsCorrect(c) {
		t.Fatalf("Expected %v to be higher than %v", b, a)
	}
}

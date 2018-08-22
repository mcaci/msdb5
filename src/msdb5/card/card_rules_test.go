package card

import "testing"

var firstCardWins = func(i int) bool { return i > 0 }
var firstCardLoses = func(i int) bool { return i < 0 }

func TestInClassicNumericalComparisonWithSameSeedHigherNumberWins(t *testing.T) {
	a := Card{number: 2, seed: Coin}
	b := Card{number: 4, seed: Coin}
	verifyCardComparison(t, &a, &b, firstCardLoses)
}

func TestComparisonWithSameSeedThat3isHigherThan10(t *testing.T) {
	a := Card{number: 10, seed: Coin}
	b := Card{number: 3, seed: Coin}
	verifyCardComparison(t, &a, &b, firstCardLoses)
}

func TestComparisonWithSameSeedThat1isHigherThan9(t *testing.T) {
	a := Card{number: 1, seed: Coin}
	b := Card{number: 9, seed: Coin}
	verifyCardComparison(t, &a, &b, firstCardWins)
}

func TestComparisonWithSameSeedThat8isHigherThan7(t *testing.T) {
	a := Card{number: 8, seed: Coin}
	b := Card{number: 7, seed: Coin}
	verifyCardComparison(t, &a, &b, firstCardWins)
}

func TestComparisonWithDifferentSeedThatFirstCardAlwaysWins(t *testing.T) {
	a := Card{number: 8, seed: Coin}
	b := Card{number: 9, seed: Cup}
	verifyCardComparison(t, &a, &b, firstCardWins)
}

func TestComparisonWithEmptyCardThatNonEmptyCardWins(t *testing.T) {
	a := Card{number: 8, seed: Coin}
	var b Card
	verifyCardComparison(t, &a, &b, firstCardWins)
}

func verifyCardComparison(t *testing.T, a, b *Card, isComparisonBetweenCardsCorrect func(int) bool) {
	c := a.Compare(*b)
	if !isComparisonBetweenCardsCorrect(c) {
		t.Fatalf("Expected %v to be higher than %v", b, a)
	}
}

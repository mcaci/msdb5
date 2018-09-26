package card

import "testing"

var firstCardWins = func(b bool) bool { return !b }
var firstCardLoses = func(b bool) bool { return b }

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

func TestComparisonWithSecondCardOfDifferentSeedThatFirstCardAlwaysWins(t *testing.T) {
	a := Card{number: 2, seed: Sword}
	b := Card{number: 3, seed: Cup}
	verifyCardComparison(t, &a, &b, firstCardWins)
}

func TestComparisonWithEmptyCardThatNonEmptyCardWins(t *testing.T) {
	a := Card{number: 8, seed: Coin}
	var b Card
	verifyCardComparison(t, &a, &b, firstCardWins)
}

func verifyCardComparison(t *testing.T, a, b *Card, isComparisonBetweenCardsCorrect func(bool) bool) {
	c := DoesOtherCardWin(a, b)
	if !isComparisonBetweenCardsCorrect(c) {
		t.Fatalf("Expected %v to be higher than %v", b, a)
	}
}

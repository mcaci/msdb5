package rule

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

var firstCardWins = func(b bool) bool { return !b }
var firstCardLoses = func(b bool) bool { return b }

func TestInClassicNumericalComparisonWithSameSeedHigherNumberWins(t *testing.T) {
	a, _ := card.ByName("2", "Coin")
	b, _ := card.ByName("4", "Coin")
	verifyCardComparison(t, a, b, firstCardLoses)
}

func TestComparisonWithSameSeedThat3isHigherThan10(t *testing.T) {
	a, _ := card.ByName("10", "Coin")
	b, _ := card.ByName("3", "Coin")
	verifyCardComparison(t, a, b, firstCardLoses)
}

func TestComparisonWithSameSeedThat1isHigherThan9(t *testing.T) {
	a, _ := card.ByName("1", "Coin")
	b, _ := card.ByName("9", "Coin")
	verifyCardComparison(t, a, b, firstCardWins)
}

func TestComparisonWithSameSeedThat7isHigherThan6(t *testing.T) {
	a, _ := card.ByName("7", "Coin")
	b, _ := card.ByName("6", "Coin")
	verifyCardComparison(t, a, b, firstCardWins)
}

func TestComparisonWithSameSeedThat7isHigherThan6InReverseOrder(t *testing.T) {
	a, _ := card.ByName("6", "Coin")
	b, _ := card.ByName("7", "Coin")
	verifyCardComparison(t, a, b, firstCardLoses)
}

func TestComparisonWithSecondLowerCardOfDifferentSeedThatFirstCardAlwaysWins(t *testing.T) {
	a, _ := card.ByName("2", "Sword")
	b, _ := card.ByName("3", "Cup")
	verifyCardComparison(t, a, b, firstCardWins)
}

func TestComparisonWithSecondHigherCardOfDifferentSeedThatFirstCardAlwaysWins(t *testing.T) {
	a, _ := card.ByName("5", "Cudgel")
	b, _ := card.ByName("4", "Cup")
	verifyCardComparison(t, a, b, firstCardWins)
}

func TestComparisonWithEmptyCardThatNonEmptyCardWins(t *testing.T) {
	a, _ := card.ByName("8", "Coin")
	var b card.ID
	verifyCardComparison(t, a, b, firstCardWins)
}

func verifyCardComparison(t *testing.T, a, b card.ID, isOtherCardExpectedToWin func(bool) bool) {
	result := doesOtherCardWinOnAttributes(a, b)
	if !isOtherCardExpectedToWin(result) {
		t.Fatalf("Expected %v to be higher than %v", b, a)
	}
}

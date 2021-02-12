package auction

import (
	"testing"
)

var initialValue Score
var minValue Score = 61
var maxValue Score = 120

func TestRaiseAuctionScoreFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	const currentValue = 1
	value := Update(initialValue, currentValue)
	assertPlayerScore(t, value, minValue)
}

func TestInvalidRaiseAuctionScoreFirstAssignmentShouldBeAlways61(t *testing.T) {
	const currentValue = 0
	value := Update(initialValue, currentValue)
	assertPlayerScore(t, value, minValue)
}

func TestRaiseAuctionTo65(t *testing.T) {
	const currentValue = 65
	value := Update(initialValue, currentValue)
	assertPlayerScore(t, value, currentValue)
}
func TestRaiseAuctionTo135ShouldStopAt120(t *testing.T) {
	const currentValue = 135
	value := Update(initialValue, currentValue)
	assertPlayerScore(t, value, maxValue)
}

func TestPlayerRaisingAuctionAfterAnotherWithLowerScore(t *testing.T) {
	const currentValue = 90
	value := Update(Score(94), currentValue)
	assertPlayerScore(t, value, Score(94))
}

func TestCheckAndUpdate_OK(t *testing.T) {
	value := Score(80)
	if !CheckScores(value, Score(100)) {
		t.Fatal("Unexpected check return value")
	}
}

func TestCheckAndUpdate_Fold(t *testing.T) {
	value := Score(80)
	if CheckScores(value, Score(61)) {
		t.Fatal("Unexpected check return value")
	}
}

func assertPlayerScore(t *testing.T, actualScore, expectedScore Score) {
	if actualScore != expectedScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}

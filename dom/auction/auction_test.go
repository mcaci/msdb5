package auction

import (
	"testing"
)

var initialValue Score
var minValue Score = 61
var maxValue Score = 120

func TestRaiseAuctionScoreFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	const currentValue = 1
	initialValue.Update(currentValue)
	testPlayerScore(t, initialValue, minValue)
}

func TestInvalidRaiseAuctionScoreFirstAssignmentShouldBeAlways61(t *testing.T) {
	const currentValue = 0
	initialValue.Update(currentValue)
	testPlayerScore(t, initialValue, minValue)
}

func TestRaiseAuctionTo65(t *testing.T) {
	const currentValue = 65
	initialValue.Update(currentValue)
	testPlayerScore(t, initialValue, currentValue)
}
func TestRaiseAuctionTo135ShouldStopAt120(t *testing.T) {
	const currentValue = 135
	initialValue.Update(currentValue)
	testPlayerScore(t, initialValue, maxValue)
}

func TestPlayerRaisingAuctionAfterAnotherWithLowerScore(t *testing.T) {
	value1 := Score(94)
	const value2 = 90
	value1.Update(value2)
	testPlayerScore(t, value1, value1)
}

func TestCheckAndUpdate_OK(t *testing.T) {
	value := Score(80)
	if !value.CheckWith(Score(100)) {
		t.Fatal("Unexpected check return value")
	}
}

func TestCheckAndUpdate_Fold(t *testing.T) {
	value := Score(80)
	if value.CheckWith(Score(61)) {
		t.Fatal("Unexpected check return value")
	}
}

func testPlayerScore(t *testing.T, actualScore, expectedScore Score) {
	if actualScore != expectedScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}

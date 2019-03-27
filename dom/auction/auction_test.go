package auction

import (
	"strconv"
	"testing"
)

type auctionTest struct {
	auction uint8
}

func (testObject *auctionTest) set(value uint8) {
	testObject.auction = value
}
func (testObject *auctionTest) get() uint8 {
	return testObject.auction
}

const initialValue = 0
const minValue = 61
const maxValue = 120

func TestRaiseAuctionScoreFirstAssignmentShouldBeSuperiorThan61ElseEither61(t *testing.T) {
	const currentValue = 1
	testObject := auctionTest{initialValue}
	Update(initialValue, currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, minValue)
}

func TestInvalidRaiseAuctionScoreFirstAssignmentShouldBeAlways61(t *testing.T) {
	const currentValue = 0
	testObject := auctionTest{initialValue}
	Update(initialValue, currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, minValue)
}

func TestRaiseAuctionTo65(t *testing.T) {
	const currentValue = 65
	testObject := auctionTest{initialValue}
	Update(initialValue, currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, currentValue)
}
func TestRaiseAuctionTo135ShouldStopAt120(t *testing.T) {
	const currentValue = 135
	testObject := auctionTest{initialValue}
	Update(initialValue, currentValue, testObject.set)
	testPlayerScore(t, testObject.auction, maxValue)
}

func TestPlayerRaisingAuctionAfterAnotherWithLowerScore(t *testing.T) {
	const value1 = 94
	testObject := auctionTest{value1}
	const value2 = 90
	Update(value1, value2, testObject.set)
	testPlayerScore(t, testObject.auction, value1)
}

func Test2PlayersRaisingAuction(t *testing.T) {
	const value1 = 65
	const value2 = 80
	testObject := auctionTest{initialValue}
	Update(initialValue, value1, testObject.set)
	Update(value1, value2, testObject.set)
	testPlayerScore(t, testObject.auction, value2)
}

func TestCheckAndUpdate_OK(t *testing.T) {
	const value = 80
	testObject := auctionTest{initialValue}
	CheckAndUpdate(strconv.Itoa(value), func() bool { return false }, func() {}, testObject.get, testObject.set)
	testPlayerScore(t, testObject.auction, value)
}

func TestCheckAndUpdate_Fold(t *testing.T) {
	testObject := auctionTest{initialValue}
	CheckAndUpdate("ciao", func() bool { return false }, func() {}, testObject.get, testObject.set)
	testPlayerScore(t, testObject.auction, initialValue)
}

func testPlayerScore(t *testing.T, actualScore, expectedScore uint8) {
	if actualScore != expectedScore {
		t.Fatalf("Auction score should be set at %d but is %d", expectedScore, actualScore)
	}
}

package card

import (
	"strconv"
	"testing"
)

func Test1OfCoinIsCreatedCorrectly_NoError(t *testing.T) {
	NoErrorCheck(t, "1", "Coin")
}

func Test1OfCoinIsCreatedCorrectly_NumberIs1(t *testing.T) {
	NumberOfCardCheck(t, "1", "Coin")
}

func Test1OfCoinIsCreatedCorrectly_SeedIsCoin(t *testing.T) {
	SeedOfCardCheck(t, "1", "Coin")
}

func Test2OfSwordIsCreatedCorrectly_NoError(t *testing.T) {
	NoErrorCheck(t, "2", "Sword")
}

func Test2OfSwordIsCreatedCorrectly_NumberIs2(t *testing.T) {
	NumberOfCardCheck(t, "2", "Sword")
}

func Test2OfSwordIsCreatedCorrectly_SeedIsSword(t *testing.T) {
	SeedOfCardCheck(t, "2", "Sword")
}

func Test8OfCupIsCreatedCorrectly_NoError(t *testing.T) {
	NoErrorCheck(t, "8", "Cup")
}

func Test8OfCupIsCreatedCorrectly_NumberIs8(t *testing.T) {
	NumberOfCardCheck(t, "8", "Cup")
}

func Test8OfCupIsCreatedCorrectly_SeedIsCup(t *testing.T) {
	SeedOfCardCheck(t, "8", "Cup")
}

func Test10OfCudgelIsCreatedCorrectly_NoError(t *testing.T) {
	NoErrorCheck(t, "10", "Cudgel")
}

func Test10OfCudgelIsCreatedCorrectly_NumberIs10(t *testing.T) {
	NumberOfCardCheck(t, "10", "Cudgel")
}

func Test10OfCudgelIsCreatedCorrectly_SeedIsCudgel(t *testing.T) {
	SeedOfCardCheck(t, "10", "Cudgel")
}

func Test15OfCupDoesntExist(t *testing.T) {
	errorCheck(t, "15", "Cup")
}

func Test8OfSpadesDoesntExist(t *testing.T) {
	errorCheck(t, "8", "Spades")
}

func TestTwoOfCudgelIsIncorrect(t *testing.T) {
	errorCheck(t, "Two", "Cudgel")
}

func TestEmptyNumberIsIncorrect(t *testing.T) {
	errorCheck(t, "", "Cudgel")
}

func TestEmptySeedIsIncorrect(t *testing.T) {
	errorCheck(t, "6", "")
}

func SeedOfCardCheck(t *testing.T, number, seed string) {
	check := func(card Data, err error) bool { return card.Seed().String() != seed }
	if check(ByName(number, seed)) {
		t.Fatalf("Card's number is not created well from %s and %s", number, seed)
	}
}

func NumberOfCardCheck(t *testing.T, number, seed string) {
	check := func(card Data, err error) bool { return strconv.Itoa(int(card.Number())) != number }
	if check(ByName(number, seed)) {
		t.Fatalf("Card's number is not created well from %s and %s", number, seed)
	}
}

func NoErrorCheck(t *testing.T, number, seed string) {
	check := func(card Data, err error) bool { return err != nil }
	if check(ByName(number, seed)) {
		t.Fatalf("An unexpected error was raised")
	}
}

func errorCheck(t *testing.T, number, seed string) {
	check := func(card Data, err error) bool { return err == nil }
	if check(ByName(number, seed)) {
		t.Fatalf("The %s of %s isn't a valid card", number, seed)
	}
}

package card

import (
	"strconv"
	"testing"
)

func NoErrorCheck(t *testing.T, number, seed string) {
	_, err := ByName(number, seed)
	if err != nil {
		t.Fatal("An unexpected error was raised")
	}
}

func NumberOfCardCheck(t *testing.T, number, seed string) {
	card, _ := ByName(number, seed)
	if strconv.Itoa(int(card.Number())) != number {
		t.Fatalf("Card %v's number is not created well from %s and %s", card, number, seed)
	}
}

func SeedOfCardCheck(t *testing.T, number, seed string) {
	card, _ := ByName(number, seed)
	if card.Seed().String() != seed {
		t.Fatalf("Card %v's number is not created well from %s and %s", card, number, seed)
	}
}

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

func errorCheck(t *testing.T, number, seed string) {
	_, err := ByName(number, seed)
	if err == nil {
		t.Fatal("The " + number + " of " + seed + " isn't a valid card")
	}
}

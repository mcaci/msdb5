package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func TestBriscolaScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin, briscola is Coin
	b := card.Coin
	verifySortingWithBriscola(t, 1, 2, &b)
}

func TestBriscolaScenarioWithTwoOfCoinLosing(t *testing.T) {
	// 2 and 3 of Coin, briscola is Coin
	b := card.Coin
	verifySortingWithBriscola(t, 3, 2, &b)
}

func TestBriscolaScenarioWithSixOfCoinWinningBecauseHigher(t *testing.T) {
	// 5 and 6 of Coin, briscola is Coin
	b := card.Coin
	verifySortingWithBriscola(t, 6, 5, &b)
}

func TestBriscolaScenarioWithTenOfCoinWinning(t *testing.T) {
	// 10 and 4 of Coin, briscola is Cup
	b := card.Cup
	verifySortingWithBriscola(t, 10, 4, &b)
}

func TestBriscolaScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
	// 3 of Coin and 2 of Sword, briscola is Sword
	b := card.Sword
	verifySortingWithBriscola(t, 22, 3, &b)
}

func verifySortingWithBriscola(t *testing.T, a, b uint8, br *card.Seed) {
	s := SortedCard{*set.NewMust(a, b), br}
	if s.Less(0, 1) {
		t.Fatalf("Expecting %d to be bigger than %d with briscola %v", a, b, br)
	}
}

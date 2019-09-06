package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func TestSortScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin, briscola is Coin
	b := card.Coin
	s := sortedCard{*set.NewMust(1, 2), &b}
	if s.Less(0, 1) {
		t.Fatal("Expecting 1 of Coin to be bigger")
	}
}

// func TestSortScenarioWithTwoOfCoinLosing(t *testing.T) {
// 	// 2 and 3 of Coin, briscola is Coin
// 	verifyRoundScenario(t, 2, 3, testSeeder(card.Coin), true)
// }

// func TestSortScenarioWithSixOfCoinWinningBecauseHigher(t *testing.T) {
// 	// 5 and 6 of Coin, briscola is Coin
// 	verifyRoundScenario(t, 5, 6, testSeeder(card.Coin), true)
// }

// func TestSortScenarioWithSixOfCoinWinningBecausePlayedFirst(t *testing.T) {
// 	// 6 and 5 of Coin, briscola is Coin
// 	verifyRoundScenario(t, 6, 5, testSeeder(card.Coin), false)
// }

// func TestSortScenarioWithTenOfCoinWinning(t *testing.T) {
// 	// 10 and 4 of Coin, briscola is Cup
// 	verifyRoundScenario(t, 10, 4, testSeeder(card.Cup), false)
// }

// func TestSortScenarioWithTenOfCoinLosing(t *testing.T) {
// 	// 10 and 4 of Coin, briscola is Cup
// 	verifyRoundScenario(t, 10, 3, testSeeder(card.Coin), true)
// }

// func TestSortScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
// 	// 3 of Coin and 2 of Sword, briscola is Sword
// 	verifyRoundScenario(t, 3, 22, testSeeder(card.Sword), true)
// }

// func verifyRoundScenario(t *testing.T, a, b uint8, briscola testSeeder, expectedWinner bool) {
// 	if index := doesOtherCardWin(*card.MustID(a), *card.MustID(b), briscola); index != expectedWinner {
// 		t.Fatal("Unexpected winner")
// 	}
// }

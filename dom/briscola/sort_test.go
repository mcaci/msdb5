package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
)

func TestScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin, briscola is Coin
	s := sortedCard{*set.NewMust(1, 2), nil}
	if s.Less(0, 1) {
		t.Fatal("Expecting 1 of Coin to be bigger")
	}
}

// func TestScenarioWithTwoOfCoinLosing(t *testing.T) {
// 	// 2 and 3 of Coin, briscola is Coin
// 	verifyRoundScenario(t, 2, 3, card.Coin, true)
// }

// func TestScenarioWithSixOfCoinWinningBecauseHigher(t *testing.T) {
// 	// 5 and 6 of Coin, briscola is Coin
// 	verifyRoundScenario(t, 5, 6, card.Coin, true)
// }

// func TestScenarioWithSixOfCoinWinningBecausePlayedFirst(t *testing.T) {
// 	// 6 and 5 of Coin, briscola is Coin
// 	verifyRoundScenario(t, 6, 5, card.Coin, false)
// }

// func TestScenarioWithTenOfCoinWinning(t *testing.T) {
// 	// 10 and 4 of Coin, briscola is Cup
// 	verifyRoundScenario(t, 10, 4, card.Cup, false)
// }

// func TestScenarioWithTenOfCoinLosing(t *testing.T) {
// 	// 10 and 4 of Coin, briscola is Cup
// 	verifyRoundScenario(t, 10, 3, card.Coin, true)
// }

// func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
// 	// 3 of Coin and 2 of Sword, briscola is Sword
// 	verifyRoundScenario(t, 3, 22, card.Sword, true)
// }

// func verifyRoundScenario(t *testing.T, a, b uint8, briscola card.Seed, expectedWinner bool) {
// 	if index := doesOtherCardWin(*card.MustID(a), *card.MustID(b), briscola); index != expectedWinner {
// 		t.Fatal("Unexpected winner")
// 	}
// }

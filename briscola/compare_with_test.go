package briscola

import "testing"
import "github.com/nikiforosFreespirit/msdb5/card"

func TestScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin, briscola is Coin
	verifyRoundScenario(t, 1, 2, card.Coin, false)
}

func TestScenarioWithTwoOfCoinLosing(t *testing.T) {
	// 2 and 3 of Coin, briscola is Coin
	verifyRoundScenario(t, 2, 3, card.Coin, true)
}

func TestScenarioWithSixOfCoinWinningBecauseHigher(t *testing.T) {
	// 5 and 6 of Coin, briscola is Coin
	verifyRoundScenario(t, 5, 6, card.Coin, true)
}

func TestScenarioWithSixOfCoinWinningBecausePlayedFirst(t *testing.T) {
	// 6 and 5 of Coin, briscola is Coin
	verifyRoundScenario(t, 6, 5, card.Coin, false)
}

func TestScenarioWithTenOfCoinWinning(t *testing.T) {
	// 10 and 4 of Coin, briscola is Cup
	verifyRoundScenario(t, 10, 4, card.Cup, false)
}

func TestScenarioWithTenOfCoinLosing(t *testing.T) {
	// 10 and 4 of Coin, briscola is Cup
	verifyRoundScenario(t, 10, 3, card.Coin, true)
}

func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
	// 3 of Coin and 2 of Sword, briscola is Sword
	verifyRoundScenario(t, 3, 22, card.Sword, true)
}

func verifyRoundScenario(t *testing.T, a, b card.ID, briscola card.Seed, expectedWinner bool) {
	if index := doesOtherCardWin(a, b, briscola); index != expectedWinner {
		t.Fatal("Unexpected winner")
	}
}

package rule

import "testing"
import "github.com/nikiforosFreespirit/msdb5/card"

func TestScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 and 2 of Coin, briscola is Coin
	verifyRoundScenario(t, 1, 2, card.Coin, false)
}

func TestScenarioWithThreeOfCoinWinning(t *testing.T) {
	// 2 and 3 of Coin, briscola is Coin
	verifyRoundScenario(t, 2, 3, card.Coin, true)
}

func TestScenarioWithEightOfCoinWinning(t *testing.T) {
	// 10 and 4 of Coin, briscola is Cup
	verifyRoundScenario(t, 10, 4, card.Cup, false)
}

func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
	// 3 of Coin and 2 of Sword, briscola is Sword
	verifyRoundScenario(t, 3, 22, card.Sword, true)
}

func verifyRoundScenario(t *testing.T, a, b card.ID, briscola card.Seed, expectedWinner bool) {
	if index := DoesOtherCardWin(a, b, briscola); index != expectedWinner {
		t.Fatal("Unexpected winner")
	}
}

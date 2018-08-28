package round

import "testing"
import "msdb5/card"

func verifyRoundScenario(t *testing.T, expectedWinner uint8, first, second, third, fourth, fifth *card.Card, briscola card.Seed) {
	i := declareWinner(first, second, third, fourth, fifth, briscola)
	if i != expectedWinner {
		t.Fatalf("Unexpected winner: winner was %d", i)
	}
}

func TestScenario1WithAceOfCoinWinning(t *testing.T) {
	first, _ := card.ByID(1)  // 1 Coin
	second, _ := card.ByID(2) // 2 Coin
	third, _ := card.ByID(3)  // 3 Coin
	fourth, _ := card.ByID(4) // 4 Coin
	fifth, _ := card.ByID(5)  // 5 Coin
	briscola := card.Coin
	verifyRoundScenario(t, 0, first, second, third, fourth, fifth, briscola)
}

func TestScenario1WithThreeOfCoinWinning(t *testing.T) {
	first, _ := card.ByID(2)  // 2 Coin
	second, _ := card.ByID(3) // 3 Coin
	third, _ := card.ByID(4)  // 4 Coin
	fourth, _ := card.ByID(5) // 5 Coin
	fifth, _ := card.ByID(6)  // 6 Coin
	briscola := card.Coin
	verifyRoundScenario(t, 1, first, second, third, fourth, fifth, briscola)
}

func TestScenario1WithEightOfCoinWinning(t *testing.T) {
	first, _ := card.ByID(4)  // 4 Coin
	second, _ := card.ByID(5) // 5 Coin
	third, _ := card.ByID(6)  // 6 Coin
	fourth, _ := card.ByID(7) // 7 Coin
	fifth, _ := card.ByID(8)  // 8 Coin
	briscola := card.Coin
	verifyRoundScenario(t, 4, first, second, third, fourth, fifth, briscola)
}

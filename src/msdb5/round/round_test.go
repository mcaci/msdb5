package round

import "testing"
import "msdb5/card"

func TestScenario1WithAceOfCoinWinning(t *testing.T) {
	verifyRoundScenario(t, 0, cardsOnTheTable(1, 2, 3, 4, 5), card.Coin)
}

func TestScenario1WithThreeOfCoinWinning(t *testing.T) {
	verifyRoundScenario(t, 1, cardsOnTheTable(2, 3, 4, 5, 6), card.Coin)
}

func TestScenario1WithEightOfCoinWinning(t *testing.T) {
	verifyRoundScenario(t, 4, cardsOnTheTable(4, 5, 6, 7, 8), card.Coin)
}

func cardsOnTheTable(cardIds ...uint8) [5]*card.Card {
	var cards [5]*card.Card
	for i := range cards {
		cards[i], _ = card.ByID(cardIds[i])
	}
	return cards
}

func verifyRoundScenario(t *testing.T, expectedWinner uint8, cardsOnTheTable [5]*card.Card, briscola card.Seed) {
	i := declareWinner(cardsOnTheTable, briscola)
	if i != expectedWinner {
		t.Fatalf("Unexpected winner: winner was %d", i)
	}
}
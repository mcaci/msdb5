package round

import "testing"
import "msdb5/card"

func TestScenario1WithAceOfCoinWinning(t *testing.T) {
	verifyRoundScenario(t, cardsOnTheTable(1, 2, 3, 4, 5), card.Coin, 0)
}

func TestScenario1WithThreeOfCoinWinning(t *testing.T) {
	verifyRoundScenario(t, cardsOnTheTable(2, 3, 4, 5, 6), card.Coin, 1)
}

func TestScenario1WithEightOfCoinWinning(t *testing.T) {
	verifyRoundScenario(t, cardsOnTheTable(4, 5, 6, 7, 8), card.Coin, 4)
}

func cardsOnTheTable(cardIds ...uint8) [5]*card.Card {
	var cards [5]*card.Card
	for i := range cards {
		cards[i], _ = card.ByID(cardIds[i])
	}
	return cards
}

func verifyRoundScenario(t *testing.T, cardsOnTheTable [5]*card.Card, briscola card.Seed, expectedWinner uint8) {
	index := IndexOfWinningCard(cardsOnTheTable, briscola)
	if index != expectedWinner {
		t.Fatalf("Unexpected winner: winner was %d", index)
	}
}

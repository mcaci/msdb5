package round

import "testing"
import "msdb5/card"

func TestScenario1WithAceOfCoinWinning(t *testing.T) {
	// testing 1 to 5 of Coin
	verifyRoundScenario(t, cardsOnTheTable(1, 2, 3, 4, 5), card.Coin, 0)
}

func TestScenario1WithThreeOfCoinWinning(t *testing.T) {
	// 2 to 6 of Coin
	verifyRoundScenario(t, cardsOnTheTable(2, 3, 4, 5, 6), card.Coin, 1)
}

func TestScenario1WithEightOfCoinWinning(t *testing.T) {
	// 4 to 8 of Coin
	verifyRoundScenario(t, cardsOnTheTable(4, 5, 6, 7, 8), card.Coin, 4)
}

func TestScenario1WithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyRoundScenario(t, cardsOnTheTable(34, 40, 3, 22, 15), card.Sword, 3)
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

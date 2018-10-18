package round

import "testing"
import "github.com/nikiforosFreespirit/msdb5/card"

func TestScenarioWithAceOfCoinWinning(t *testing.T) {
	// testing 1 to 5 of Coin
	verifyRoundScenario(t, card.Fill(card.WithIDs, 1, 2, 3, 4, 5), card.Coin, 0)
}

func TestScenarioWithThreeOfCoinWinning(t *testing.T) {
	// 2 to 6 of Coin
	verifyRoundScenario(t, card.Fill(card.WithIDs, 2, 3, 4, 5, 6), card.Coin, 1)
}

func TestScenarioWithEightOfCoinWinning(t *testing.T) {
	// 4 to 8 of Coin
	verifyRoundScenario(t, card.Fill(card.WithIDs, 4, 5, 6, 7, 8), card.Cup, 4)
}

func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscola(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyRoundScenario(t, card.Fill(card.WithIDs, 34, 40, 3, 22, 15), card.Sword, 3)
}

func TestScenarioWithTwoBriscolaCards(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyRoundScenario(t, card.Fill(card.WithIDs, 34, 40, 3, 22, 15), card.Cudgel, 1)
}
func TestScenarioWithTwoBriscolaCardsAndHighCardAtTheEnd(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 1 of Cup
	verifyRoundScenario(t, card.Fill(card.WithIDs, 34, 40, 3, 22, 11), card.Cudgel, 1)
}

func verifyRoundScenario(t *testing.T, cardsOnTheTable card.Cards, briscola card.Seed, expectedWinner uint8) {
	index := IndexOfWinningCard(cardsOnTheTable, briscola)
	if index != expectedWinner {
		t.Fatalf("Unexpected winner: winner was %d", index)
	}
}

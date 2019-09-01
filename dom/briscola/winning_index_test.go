package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
)

func TestScenarioWithAceOfCoinWinningIndex(t *testing.T) {
	// testing 1 to 5 of Coin
	verifyIndexScenario(t, set.NewMust(1, 2, 3, 4, 5), testSeeder(card.Coin), 0)
}

func TestScenarioWithThreeOfCoinWinningIndex(t *testing.T) {
	// 2 to 6 of Coin
	verifyIndexScenario(t, set.NewMust(2, 3, 4, 5, 6), testSeeder(card.Coin), 1)
}

func TestScenarioWithEightOfCoinWinningIndex(t *testing.T) {
	// 4 to 8 of Coin
	verifyIndexScenario(t, set.NewMust(4, 5, 6, 7, 8), testSeeder(card.Cup), 4)
}

func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscolaIndex(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyIndexScenario(t, set.NewMust(34, 40, 3, 22, 15), testSeeder(card.Sword), 3)
}

func TestScenarioWithTwoBriscolaCardsIndex(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyIndexScenario(t, set.NewMust(34, 40, 3, 22, 15), testSeeder(card.Cudgel), 1)
}
func TestScenarioWithTwoBriscolaCardsAndHighCardAtTheEndIndex(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 9 of Coin, 2 of Sword, 1 of Cup
	verifyIndexScenario(t, set.NewMust(34, 40, 9, 22, 11), testSeeder(card.Cudgel), 1)
}

func verifyIndexScenario(t *testing.T, cardsOnTheTable *set.Cards, briscola testSeeder, expectedWinner uint8) {
	if index := IndexOfWinningCard(*cardsOnTheTable, briscola); index != expectedWinner {
		t.Fatalf("Unexpected winner: winner was %d", index)
	}
}

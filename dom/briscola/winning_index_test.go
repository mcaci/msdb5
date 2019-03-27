package briscola

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

func TestScenarioWithAceOfCoinWinningIndex(t *testing.T) {
	// testing 1 to 5 of Coin
	verifyIndexScenario(t, deck.Cards{1, 2, 3, 4, 5}, card.Coin, 0)
}

func TestScenarioWithThreeOfCoinWinningIndex(t *testing.T) {
	// 2 to 6 of Coin
	verifyIndexScenario(t, deck.Cards{2, 3, 4, 5, 6}, card.Coin, 1)
}

func TestScenarioWithEightOfCoinWinningIndex(t *testing.T) {
	// 4 to 8 of Coin
	verifyIndexScenario(t, deck.Cards{4, 5, 6, 7, 8}, card.Cup, 4)
}

func TestScenarioWithTwoOfSwordsWinningBecauseOfBriscolaIndex(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyIndexScenario(t, deck.Cards{34, 40, 3, 22, 15}, card.Sword, 3)
}

func TestScenarioWithTwoBriscolaCardsIndex(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 3 of Coin, 2 of Sword, 5 of Cup
	verifyIndexScenario(t, deck.Cards{34, 40, 3, 22, 15}, card.Cudgel, 1)
}
func TestScenarioWithTwoBriscolaCardsAndHighCardAtTheEndIndex(t *testing.T) {
	// 4 of Cudgel, 10 of Cudgel, 9 of Coin, 2 of Sword, 1 of Cup
	verifyIndexScenario(t, deck.Cards{34, 40, 9, 22, 11}, card.Cudgel, 1)
}

func verifyIndexScenario(t *testing.T, cardsOnTheTable deck.Cards, briscola card.Seed, expectedWinner uint8) {
	if index := IndexOfWinningCard(cardsOnTheTable, briscola); index != expectedWinner {
		t.Fatalf("Unexpected winner: winner was %d", index)
	}
}

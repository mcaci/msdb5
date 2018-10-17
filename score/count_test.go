package score

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
	"github.com/nikiforosFreespirit/msdb5/deck"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards ...card.Card) {
	score := Compute(cards...)
	if expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func deckCards() []card.Card {
	var cards []card.Card
	deck := deck.New()
	for !deck.IsEmpty() {
		cards = append(cards, deck.Supply())
	}
	return cards
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	ace, _ := card.ByID(1)
	testScoreCount(t, 11, ace)
}

func TestPileWithOneTwoOnehAceOnlySums11(t *testing.T) {
	ace, _ := card.ByID(1)
	two, _ := card.ByID(2)
	testScoreCount(t, 11, two, ace)
}

func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, deckCards()...)
}

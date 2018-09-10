package score

import (
	"msdb5/card"
	"msdb5/deck"
	"testing"
)

func testScoreCount(t *testing.T, expectedScore uint8, cards ...*card.Card) {
	score := Compute(cards...)
	if expectedScore != score {
		t.Fatalf("Score expected is %d but %d was computed", expectedScore, score)
	}
}

func deckCards() []*card.Card {
	var cards []*card.Card
	deck := deck.New()
	for !deck.IsEmpty() {
		cards = append(cards, deck.RemoveTop())
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

func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, deckCards()...)
}

func TestPileWithOneTwoOnehAceOnlySums11(t *testing.T) {
	two, _ := card.ByID(2)
	ace, _ := card.ByID(1)
	testScoreCount(t, 11, two, ace)
}

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

func cards(fill func(...uint8) card.Cards, ids ...uint8) card.Cards {
	return fill(ids...)
}

func fromDeck(ids ...uint8) card.Cards {
	var cards card.Cards
	deck := deck.New()
	cs := deck.Get()
	i := 0
	for i < len(cs) {
		card, _ := card.ByID(uint8(cs[i] + 1))
		cards.Add(card)
		i++
	}
	return cards
}

func withIDs(ids ...uint8) card.Cards {
	var cards card.Cards
	i := 0
	for i < len(ids) {
		card, _ := card.ByID(ids[i])
		cards.Add(card)
		i++
	}
	return cards
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, cards(withIDs, 1)...)
}

func TestPileWithOneTwoOneAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, cards(withIDs, 2, 1)...)
}

func TestPileWithOneAceOneTwoOneThreeSums21(t *testing.T) {
	testScoreCount(t, 21, cards(withIDs, 1, 2, 3)...)
}
func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, cards(fromDeck)...)
}

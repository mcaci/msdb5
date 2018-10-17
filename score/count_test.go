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

func cards(ids ...uint8) card.Cards {
	var cards card.Cards
	if len(ids) == 0 {
		deck := deck.New()
		for !deck.IsEmpty() {
			cards = append(cards, deck.Supply())
		}
	} else {
		for _, id := range ids {
			card, _ := card.ByID(id)
			cards = append(cards, card)
		}
	}
	return cards
}

func TestEmptyPileSums0(t *testing.T) {
	testScoreCount(t, 0)
}

func TestPileWithOnehAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, cards(1)...)
}

func TestPileWithOneTwoOneAceOnlySums11(t *testing.T) {
	testScoreCount(t, 11, cards(2, 1)...)
}

func TestPileWithOneAceOneTwoOneThreeSums21(t *testing.T) {
	testScoreCount(t, 21, cards(1, 2, 3)...)
}
func TestPileWithAllCardsSums120(t *testing.T) {
	testScoreCount(t, 120, cards()...)
}

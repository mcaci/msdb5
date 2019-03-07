package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestPlayerPlaysCard(t *testing.T) {
	if initTest("1", "Coin").err != nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeChangesIfPlayerPlaysCardInHand(t *testing.T) {
	if initTest("1", "Coin").diffHandLenght != 1 {
		t.Fatal("Hand before playing owned card should contain one more card")
	}
}

func TestErrIfPlayerPlaysCardNotInHand(t *testing.T) {
	if initTest("2", "Coin").err == nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestNoCardReturnedIfPlayerPlaysCardNotInHand(t *testing.T) {
	if initTest("2", "Coin").playedCard != 0 {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeDoesntChangeIfPlayerPlaysCardNotInHand(t *testing.T) {
	if initTest("2", "Coin").diffHandLenght != 0 {
		t.Fatal("In case of error handsize should not change")
	}
}

type dataTest struct {
	diffHandLenght int
	playedCard     card.ID
	err            error
}

func initTest(number, seed string) dataTest {
	p := New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	cardID, err := p.Play(number, seed)
	return dataTest{len(oldHand) - len(p.hand), cardID, err}
}

package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/card"
)

func TestPlayerPlaysCard(t *testing.T) {
	if testPlay("1", "Coin").err != nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeChangesIfPlayerPlaysCardInHand(t *testing.T) {
	if testPlay("1", "Coin").diffHandLenght != 1 {
		t.Fatal("Hand before playing owned card should contain one more card")
	}
}

func TestErrIfPlayerPlaysCardNotInHand(t *testing.T) {
	if testPlay("2", "Coin").err == nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeDoesntChangeIfPlayerPlaysCardNotInHand(t *testing.T) {
	if testPlay("2", "Coin").diffHandLenght != 0 {
		t.Fatal("In case of error handsize should not change")
	}
}

type dataTest struct {
	diffHandLenght int
	err            error
}

func testPlay(number, seed string) dataTest {
	p := New()
	cardID := card.ID(1)
	p.Draw(func() card.ID { return cardID })
	oldHand := p.hand
	c, err := card.Create(number, seed)
	err = p.Play(c)
	return dataTest{len(oldHand) - len(p.hand), err}
}

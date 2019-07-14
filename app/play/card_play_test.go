package play

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
)

func TestPlayerPlaysCard(t *testing.T) {
	if fakePlay("1", "Coin").err != nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeChangesIfPlayerPlaysCardInHand(t *testing.T) {
	if fakePlay("1", "Coin").diffHandLenght != 1 {
		t.Fatal("Hand before playing owned card should contain one more card")
	}
}

func TestErrIfPlayerPlaysCardNotInHand(t *testing.T) {
	if fakePlay("2", "Coin").err == nil {
		t.Fatal("Card should come from player's hand")
	}
}

func TestHandSizeDoesntChangeIfPlayerPlaysCardNotInHand(t *testing.T) {
	if fakePlay("2", "Coin").diffHandLenght != 0 {
		t.Fatal("In case of error handsize should not change")
	}
}

type dataTest struct {
	diffHandLenght int
	err            error
}

func fakePlay(number, seed string) dataTest {
	p := player.New()
	p.Draw(func() card.ID { return 1 })
	oldHand := *p.Hand()
	c, err := card.Create(number, seed)
	err = Play(c, p.Hand())
	return dataTest{len(oldHand) - len(*p.Hand()), err}
}

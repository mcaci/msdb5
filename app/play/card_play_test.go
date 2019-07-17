package play

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
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

type fakeInput card.ID

func (rq fakeInput) Card() (card.ID, error) {
	return card.ID(rq), nil
}

func fakePlay(number, seed string) dataTest {
	p := player.New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	c, _ := card.Create(number, seed)
	err := CardAction(fakeInput(c), p.Hand(), &deck.Cards{}, func(cards, to *deck.Cards, index, toIndex int) {
		to.Add((*cards)[index])
		*cards = append((*cards)[:index], (*cards)[index+1:]...)
	})
	return dataTest{len(oldHand) - len(*p.Hand()), err}
}

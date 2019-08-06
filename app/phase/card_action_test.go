package phase

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/player"
)

type cardactiontest string

func (c cardactiontest) Find(player.Predicate) (int, *player.Player) {
	return 1, nil
}

type fakeInput card.ID

func (rq fakeInput) Card() (card.ID, error) {
	return card.ID(rq), nil
}

func (c cardactiontest) Value() string { return string(c) }

func TestCardActionIndex(t *testing.T) {
	data := CardAction(fakeInput(1), cardactiontest("A"))
	if data.Index() != 1 {
		t.Fatal("Unexpected player")
	}
}

func TestCardActionCard(t *testing.T) {
	data := CardAction(fakeInput(11), cardactiontest("A"))
	if data.Card() != 11 {
		t.Fatal("Unexpected briscola")
	}
}

type errortest struct{}

func (e errortest) Find(player.Predicate) (int, *player.Player) {
	return -1, nil
}

func TestCardActionErr(t *testing.T) {
	data := CardAction(fakeInput(11), errortest{})
	if data.CardNotFound() != nil {
		t.Fatal("Error is expected")
	}
}

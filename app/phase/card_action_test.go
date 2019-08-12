package phase

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type cardactiontest string

func (c cardactiontest) Find(player.Predicate) (int, *player.Player) {
	return 0, &player.Player{}
}
func (c cardactiontest) Value() string { return string(c) }

type fakeInput struct{ c *card.Item }

func (rq fakeInput) Card() (*card.Item, error) {
	return rq.c, nil
}

func TestCardActionIndex(t *testing.T) {
	data := CardAction(fakeInput{card.MustID(1)}, cardactiontest("A"))
	if data.Pl() == nil {
		t.Fatal("Unexpected player")
	}
}

func TestCardActionCard(t *testing.T) {
	data := CardAction(fakeInput{card.MustID(11)}, cardactiontest("A"))
	if *data.Card() != *card.MustID(11) {
		t.Fatalf("Unexpected briscola, found %v", *data.Card())
	}
}

type errortest struct{}

func (e errortest) Find(player.Predicate) (int, *player.Player) {
	return -1, nil
}

func TestCardActionErr(t *testing.T) {
	data := CardAction(fakeInput{card.MustID(11)}, errortest{})
	if data.CardErr() == nil {
		t.Fatal("Error is expected")
	}
}

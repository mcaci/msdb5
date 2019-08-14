package action

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/dom/player"
)

type cardactiontest struct{}

func (cardactiontest) Find(player.Predicate) (int, *player.Player) { return 0, &player.Player{} }

type fakeCardValueProv struct {
	c   *card.Item
	str string
}

func (rq fakeCardValueProv) Card() (*card.Item, error) { return rq.c, nil }
func (rq fakeCardValueProv) Value() string             { return string(rq.str) }

type actionertest struct{}

func (actionertest) exec(plCProv playerCardProvider) {}
func (actionertest) notAcceptedZeroErr() error       { return nil }

func TestCardActionOk(t *testing.T) {
	err := CardAction(fakeCardValueProv{card.MustID(11), "A"}, cardactiontest{}, actionertest{})
	if err != nil {
		t.Fatal("Error is not expected")
	}
}

type errortest struct{}

func (e errortest) Find(player.Predicate) (int, *player.Player) { return -1, nil }

func TestCardActionErr(t *testing.T) {
	err := CardAction(fakeCardValueProv{card.MustID(11), "A"}, errortest{}, actionertest{})
	if err == nil {
		t.Fatal("Error is expected")
	}
}

package cardaction

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

type fakePlCardProv struct {
	pl *player.Player
}

func (fakePlCardProv) Card() *card.Item     { return card.MustID(3) }
func (f fakePlCardProv) Pl() *player.Player { return f.pl }

func TestPostExchangeCardsResult(t *testing.T) {
	pl := player.New()
	pl.Hand().Add(*set.NewMust(1, 2, 3)...)
	plCardProv := fakePlCardProv{pl}
	e := Exch{set.NewMust(4, 5, 6, 7, 8)}
	e.exec(plCardProv)
	if (*plCardProv.pl.Hand())[2] != *card.MustID(4) {
		t.Fatalf("Expecting %s, found %s", *card.MustID(4), (*plCardProv.pl.Hand())[2])
	}
}

func TestPostExchangeToResult(t *testing.T) {
	pl := player.New()
	pl.Hand().Add(*set.NewMust(1, 2, 3)...)
	plCardProv := fakePlCardProv{pl}
	toProv := set.NewMust(4, 5, 6, 7, 8)
	e := Exch{toProv}
	e.exec(plCardProv)
	if (*toProv)[4] != *card.MustID(3) {
		t.Fatalf("Expecting %s, found %s", *card.MustID(3), (*toProv)[4])
	}
}

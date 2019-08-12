package game

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

type fakeToProv struct {
	cards *set.Cards
}

func (f fakeToProv) PlayedCards() *set.Cards { return f.cards }
func (f fakeToProv) SideDeck() *set.Cards    { return f.cards }

func TestPostExchangeCardsResult(t *testing.T) {
	pl := player.New()
	cards := set.NewMust(1, 2, 3)
	pl.Hand().Add(*cards...)
	plCardProv := fakePlCardProv{pl}
	postExchange(plCardProv, fakeToProv{set.NewMust(4, 5, 6, 7, 8)})
	if (*plCardProv.pl.Hand())[2] != *card.MustID(4) {
		t.Fatalf("Expecting %s, found %s", *card.MustID(4), (*plCardProv.pl.Hand())[2])
	}
}

func TestPostExchangeToResult(t *testing.T) {
	pl := player.New()
	cards := set.NewMust(1, 2, 3)
	pl.Hand().Add(*cards...)
	plCardProv := fakePlCardProv{pl}
	toProv := fakeToProv{set.NewMust(4, 5, 6, 7, 8)}
	postExchange(plCardProv, toProv)
	if (*toProv.SideDeck())[4] != *card.MustID(3) {
		t.Log(toProv)
		t.Fatalf("Expecting %s, found %s", *card.MustID(3), (*toProv.SideDeck())[4])
	}
}

package game

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
	"github.com/mcaci/msdb5/dom/deck"
	"github.com/mcaci/msdb5/dom/player"
)

type fakePlCardProv struct {
	pl *player.Player
}

func (fakePlCardProv) Card() card.ID        { return 3 }
func (f fakePlCardProv) Pl() *player.Player { return f.pl }

type fakeToProv struct {
	cards *deck.Cards
}

func (f fakeToProv) PlayedCards() *deck.Cards { return f.cards }
func (f fakeToProv) SideDeck() *deck.Cards    { return f.cards }

func TestPostExchangeCardsResult(t *testing.T) {
	pl := player.New()
	cards := deck.Cards{1, 2, 3}
	pl.Hand().Add(cards...)
	plCardProv := fakePlCardProv{pl}
	postExchangeCard(plCardProv, fakeToProv{&deck.Cards{4, 5, 6, 7, 8}})
	if (*plCardProv.pl.Hand())[2] != 4 {
		t.Fatalf("Expecting %s, found %s", card.ID(4), (*plCardProv.pl.Hand())[2])
	}
}

func TestPostExchangeToResult(t *testing.T) {
	pl := player.New()
	cards := deck.Cards{1, 2, 3}
	pl.Hand().Add(cards...)
	plCardProv := fakePlCardProv{pl}
	toProv := fakeToProv{&deck.Cards{4, 5, 6, 7, 8}}
	postExchangeCard(plCardProv, toProv)
	if (*toProv.SideDeck())[4] != 3 {
		t.Log(toProv)
		t.Fatalf("Expecting %s, found %s", card.ID(3), (*toProv.SideDeck())[4])
	}
}

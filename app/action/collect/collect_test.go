package collect

import (
	"testing"

	"github.com/mcaci/msdb5/dom/team"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

func TestNoCollectFunc(t *testing.T) {
	p := player.New()
	cards := &set.Cards{*card.MustID(1), *card.MustID(2)}
	info := NewInfo(p, cards)
	Played(info)
	if len(*p.Pile()) > 0 {
		t.Fatal("Unexpected cards in hand")
	}
}

func TestCollectFunc(t *testing.T) {
	p := player.New()
	cards := &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(16), *card.MustID(32), *card.MustID(5)}
	info := NewInfo(p, cards)
	Played(info)
	if len(*p.Pile()) == 0 {
		t.Fatal("Unexpected cards not in hand")
	}
}

func TestCollectAllFuncCheckHands(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(10))
	cards := &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(16), *card.MustID(32), *card.MustID(5)}
	info := NewAllInfo(p, cards, team.Players{p})
	All(info)
	if len(*p.Hand()) > 0 {
		t.Fatal(p.Hand())
	}
}

func TestCollectAllFuncCheckSide(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(10))
	cards := &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(16), *card.MustID(32), *card.MustID(5)}
	info := NewAllInfo(p, cards, team.Players{p})
	All(info)
	if len(*cards) > 0 {
		t.Fatal(cards)
	}
}

func TestCollectAllFuncCheckPile(t *testing.T) {
	p := player.New()
	p.Hand().Add(*card.MustID(10))
	cards := &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(16), *card.MustID(32), *card.MustID(5)}
	info := NewAllInfo(p, cards, team.Players{p})
	All(info)
	if len(*p.Pile()) == 0 {
		t.Fatal(p.Pile())
	}
}

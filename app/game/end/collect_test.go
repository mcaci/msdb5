package end

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/ita-cards/set"
	"github.com/mcaci/msdb5/dom/player"
)

func TestNoCollectFunc(t *testing.T) {
	p := player.New()
	cards := &set.Cards{*card.MustID(1), *card.MustID(2)}
	info := NewCollectInfo(p, cards)
	Collect(info)
	if len(*p.Pile()) == 0 {
		t.Fatal("Unexpected cards in hand")
	}
}

func TestCollectFunc(t *testing.T) {
	p := player.New()
	cards := &set.Cards{*card.MustID(1), *card.MustID(2), *card.MustID(16), *card.MustID(32), *card.MustID(5)}
	info := NewCollectInfo(p, cards)
	Collect(info)
	if len(*p.Pile()) == 0 {
		t.Fatal("Unexpected cards not in hand")
	}
}

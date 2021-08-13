package ai

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/app/misc"
)

func TestCommonPile(t *testing.T) {
	a := misc.New(&misc.Options{For2P: true, Name: "A"})
	a.Pile().Add(*card.MustID(5), *card.MustID(16))
	b := misc.New(&misc.Options{For2P: true, Name: "B"})
	b.Pile().Add(*card.MustID(33), *card.MustID(21))

	if pile := CommonPile(misc.Players{a, b}); len(pile) != 4 {
		t.Fatal("Count should be 4")
	}
}

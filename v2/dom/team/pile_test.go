package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/player"
)

func TestCommonPile(t *testing.T) {
	var a player.B2Player
	a.RegisterAs("A")
	a.Pile().Add(*card.MustID(5), *card.MustID(16))
	var b player.B2Player
	b.RegisterAs("B")
	b.Pile().Add(*card.MustID(33), *card.MustID(21))

	if pile := CommonPile(Players{&a, &b}); len(pile) != 4 {
		t.Fatal("Count should be 4")
	}
}

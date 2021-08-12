package team

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/msdb5/v2/dom/briscola/player"
)

func TestCommonPile(t *testing.T) {
	a := player.New(&player.Options{For2P: true, Name: "A"})
	a.Pile().Add(*card.MustID(5), *card.MustID(16))
	b := player.New(&player.Options{For2P: true, Name: "B"})
	b.Pile().Add(*card.MustID(33), *card.MustID(21))

	if pile := CommonPile(Players{a, b}); len(pile) != 4 {
		t.Fatal("Count should be 4")
	}
}

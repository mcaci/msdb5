package player

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestCommonPile(t *testing.T) {
	a := New(&Options{For2P: true, Name: "A"})
	a.Pile().Add(*card.MustID(5), *card.MustID(16))
	b := New(&Options{For2P: true, Name: "B"})
	b.Pile().Add(*card.MustID(33), *card.MustID(21))

	if pile := CommonPile(Players{a, b}); len(pile) != 4 {
		t.Fatal("Count should be 4")
	}
}

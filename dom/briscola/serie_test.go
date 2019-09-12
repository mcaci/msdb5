package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestSerie(t *testing.T) {
	testResult := Serie(card.Sword)
	if testResult[1] != *card.MustID(23) {
		t.Fatal("unexpected card")
	}
}

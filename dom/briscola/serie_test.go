package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

type testObj struct{ seed card.Seed }

func (t testObj) Seed() card.Seed { return t.seed }

func TestSerie(t *testing.T) {
	testResult := Serie(testObj{card.Sword})
	if testResult[1] != *card.MustID(23) {
		t.Fatal("unexpected card")
	}
}

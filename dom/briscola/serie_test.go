package briscola

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
)

type testObj struct{ seed card.Seed }

func (t testObj) Seed() card.Seed { return t.seed }

func TestSerie(t *testing.T) {
	testResult := Serie(testObj{card.Sword})
	if testResult[1] != 23 {
		t.Fatal("unexpected card")
	}
}

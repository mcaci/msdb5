package player

import (
	"testing"

	"github.com/mcaci/msdb5/dom/card"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	player := New()
	player.Draw(func() card.ID { return 1 })
	if !player.Has(1) {
		t.Fatalf("Expecting player to have drawn %v", 1)
	}
}

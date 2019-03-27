package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/deck"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	player := New()
	cards := deck.Cards{1}
	player.Draw(cards)
	if !player.Has(1) {
		t.Fatalf("Expecting player to have drawn %v", 1)
	}
}

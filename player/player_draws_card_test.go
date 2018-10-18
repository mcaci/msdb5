package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/deck"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	d := deck.Deck()

	player := New()
	drawnCard := player.Draw(&d)
	if !player.Has(drawnCard) {
		t.Fatalf("Expecting player to have drawn %v", drawnCard)
	}
}

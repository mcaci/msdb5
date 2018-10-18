package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	d := card.Deck()

	player := New()
	drawnCard := player.Draw(&d)
	if !player.HasID(drawnCard) {
		t.Fatalf("Expecting player to have drawn %v", drawnCard)
	}
}

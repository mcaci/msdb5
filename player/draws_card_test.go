package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestPlayerDrawsOneCard(t *testing.T) {
	player := New()
	cards := deck.Cards{}
	cards.Add(1)
	drawnCard := player.Draw(cards)
	if !player.Has(drawnCard) {
		t.Fatalf("Expecting player to have drawn %v", drawnCard)
	}
}

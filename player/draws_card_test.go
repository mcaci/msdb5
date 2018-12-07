package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

type MockDeck struct{}

func (d *MockDeck) Supply() card.ID { return 1 }

func TestPlayerDrawsOneCard(t *testing.T) {
	player := New()
	drawnCard := player.Draw(&MockDeck{})
	if !player.Has(drawnCard) {
		t.Fatalf("Expecting player to have drawn %v", drawnCard)
	}
}

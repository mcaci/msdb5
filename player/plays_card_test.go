package player

import (
	"testing"
)

func TestPlayerPlaysCard(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	card := p.Supply()
	found := false
	for _, c := range oldHand {
		found = c == card
		if found {
			break
		}
	}
	if !found {
		t.Logf("Card played: %v", card)
		t.Logf("Hand before playing: %v", oldHand)
		t.Fatal("Card should come from player's hand")
	}
}

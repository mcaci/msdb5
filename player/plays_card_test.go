package player

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/card"
)

func TestPlayerPlaysCard(t *testing.T) {
	p := New()
	p.Hand().Add(1)
	oldHand := *p.Hand()
	card, _ := card.ByName("1", "Coin")
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

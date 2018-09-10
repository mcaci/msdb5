package player

import (
	"testing"
)

func TestPlayerPlaysCard(t *testing.T) {
	var d MockDeck
	player := New()
	player.Draw(&d)
	playedCard := player.Play()
	if playedCard == nil {
		t.Fatal("Expecting player to have played a card")
	}
}

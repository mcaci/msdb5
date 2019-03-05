package player

import (
	"testing"
)

func TestPlayerAuctionScoreSetting(t *testing.T) {
	var expectedScore uint8 = 71
	p := New()
	p.SetAuctionScore(71)
	actualScore := p.AuctionScore()
	if expectedScore != actualScore {
		t.Fatalf("Auction score should be %d but is %d", expectedScore, actualScore)
	}
}

func TestPlayerHasFolded(t *testing.T) {
	p := New()
	p.Fold()
	if !p.Folded() {
		t.Fatal("Player should have folded")
	}
}

func TestPlayerHasntFolded(t *testing.T) {
	p := New()
	if p.Folded() {
		t.Fatal("Player should not have folded")
	}
}

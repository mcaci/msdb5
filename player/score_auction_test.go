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

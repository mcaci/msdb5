package auction

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/dom/auction"
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

func TestAuctionDoNoErr(t *testing.T) {
	score := auction.Score(0)
	testObject := NewAuction("Auction#80", "127.0.0.3", &score)
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Auction phase: %v", err)
	}
}

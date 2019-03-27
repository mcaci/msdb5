package auction

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestAuctionDoNoErr(t *testing.T) {
	testObject := NewAuction("Auction#80", "127.0.0.3", board.New())
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Auction phase: %v", err)
	}
}

package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/board"
	"github.com/nikiforosFreespirit/msdb5/player"
	"github.com/nikiforosFreespirit/msdb5/playerset"
)

func TestAuctionWithSideDoNoErr(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Hand().Add(1)
	testObject := NewAuctionWithSide("Auction#80", "127.0.0.3",
		testPlayer, playerset.Players{testPlayer}, board.New())
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Auction phase: %v", err)
	}
}

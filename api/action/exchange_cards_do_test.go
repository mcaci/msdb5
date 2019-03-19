package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestExchangeDoNoErr(t *testing.T) {
	testObject := NewExchangeCards("Exchange#1#Coin", "127.0.0.2", nil)
	testPlayer := player.New()
	err := testObject.Do(testPlayer)
	if err == nil {
		t.Fatalf("Unexpected error when exchanging cards phase")
	}
}

package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestJoinDoNoErr(t *testing.T) {
	testObject := NewJoin("Join#Michi", "127.0.0.2")
	testPlayer := player.New()
	err := testObject.Do(testPlayer)
	if err != nil {
		t.Fatalf("Unexpected error from Join phase")
	}
}

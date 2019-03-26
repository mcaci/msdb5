package action

import (
	"testing"

	"github.com/nikiforosFreespirit/msdb5/player"
)

func TestPlayerFinderFindsPlayer(t *testing.T) {
	testPlayer := player.New()
	testPlayer.Join("A", "127.0.0.3")
	if testObject := NewPlayerFinder("127.0.0.3", testPlayer); !testObject.Find(testPlayer) {
		t.Fatalf("Unexpected player")
	}
}

func TestPlayerFinderDoesNotFindPlayerNotInTurn(t *testing.T) {
	if testObject := NewPlayerFinder("", nil); testObject.Find(player.New()) {
		t.Fatalf("Unexpected player")
	}
}

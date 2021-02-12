package track

import (
	"container/list"
	"testing"

	"github.com/mcaci/msdb5/v2/dom/player"
)

func TestPlayerTracking(t *testing.T) {
	lastPlaying := new(list.List)
	Player(lastPlaying, player.New())
	if lastPlaying.Len() != 1 {
		t.Fatal("exepecting 1 players")
	}
}

func TestTrack3PlayersKeeps2InList(t *testing.T) {
	lastPlaying := new(list.List)
	Player(lastPlaying, player.New())
	Player(lastPlaying, player.New())
	Player(lastPlaying, player.New())
	if lastPlaying.Len() > 2 {
		t.Fatal("exepecting 2 players")
	}
}

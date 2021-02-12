package track

import (
	"container/list"

	"github.com/mcaci/msdb5/v2/dom/player"
)

// Player func
func Player(lastPlaying *list.List, actingPlayer *player.Player) {
	lastPlaying.PushFront(actingPlayer)
	if lastPlaying.Len() > 2 {
		lastPlaying.Remove(lastPlaying.Back())
	}
}

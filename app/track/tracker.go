package track

import (
	"container/list"

	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Player func
func Player(lastPlaying *list.List, actingPlayer *player.Player) {
	lastPlaying.PushFront(actingPlayer)
	if lastPlaying.Len() > 2 {
		lastPlaying.Remove(lastPlaying.Back())
	}
}

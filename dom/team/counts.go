package team

import (
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Count func
func Count(players Players, predicate player.Predicate) (count uint8) {
	for _, p := range players {
		if predicate(p) {
			count++
		}
	}
	return
}

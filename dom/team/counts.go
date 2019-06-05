package team

import (
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Predicate type
type Predicate func(p *player.Player) bool

// Count func
func Count(players Players, predicate Predicate) (count uint8) {
	for _, p := range players {
		if predicate(p) {
			count++
		}
	}
	return
}

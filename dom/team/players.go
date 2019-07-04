package team

import (
	"github.com/nikiforosFreespirit/msdb5/dom/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p *player.Player) {
	*playerSet = append(*playerSet, p)
}

// Find func
func (playerSet Players) Find(predicate func(p *player.Player) bool) (int, *player.Player) {
	for i, p := range playerSet {
		if predicate(p) {
			return i, p
		}
	}
	return -1, nil
}

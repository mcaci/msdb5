package team

import (
	"github.com/mcaci/msdb5/dom/player"
)

// Players struct
type Players []*player.Player

// Add func
func (playerSet *Players) Add(p *player.Player) {
	*playerSet = append(*playerSet, p)
}

// At func
func (playerSet Players) At(i uint8) *player.Player { return playerSet[i] }

// Part func
func (playerSet Players) Part(predicate player.Predicate) (t1, t2 Players) {
	for _, p := range playerSet {
		if predicate(p) {
			t1.Add(p)
			continue
		}
		t2.Add(p)
	}
	return
}

package team

import (
	"github.com/mcaci/msdb5/v2/dom/player"
)

// Players is a slice of pointers to players
type Players []*player.Player

// Add adds a player to the team
func (playerSet *Players) Add(p *player.Player) {
	*playerSet = append(*playerSet, p)
}

// Part partition players in two groups according to a predicate
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
